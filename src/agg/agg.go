package agg

import (
	"fmt"
	"io"
	"sort"
)

type AggregateType int

const (
	Count AggregateType = iota
	Sum
	Avg
	Max
	Min
)

type Agg struct {
	Type    AggregateType
	TrimLen int
	BarLen  int
	Out     io.Writer
	Curr    string
	Counter map[string]int
	NumList []float64
}

const (
	BarWidth = 50
)

func (a *Agg) printCounter() {
	keys := make([]string, 0, len(a.Counter))

	for k := range a.Counter {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Fprint(a.Out, a.Curr)

	for _, k := range keys {
		fmt.Fprintf(a.Out, "\t%s:%d", k, a.Counter[k])
	}

	fmt.Fprintln(a.Out)
}

func (a *Agg) PushStr(ts string, value string) {
	if a.Curr != ts {
		if a.Curr != "" {
			a.printCounter()
		}

		a.Curr = ts
		a.Counter = map[string]int{}
	}

	if _, ok := a.Counter[value]; ok {
		a.Counter[value] += 1
	} else {
		a.Counter[value] = 1
	}
}

func (a *Agg) printNum() {
	var aggregated float64

	switch a.Type {
	case Sum:
		aggregated = AggregatSum(a.NumList)
	case Avg:
		aggregated = AggregatAvg(a.NumList)
	case Max:
		aggregated = AggregatMax(a.NumList)
	case Min:
		aggregated = AggregatMin(a.NumList)
	}

	fmt.Fprintf(a.Out, "%s\t%f", a.Curr, aggregated)

	if a.BarLen > 0 {
		fmt.Fprintf(a.Out, "\t%s", Bar(aggregated, BarWidth, a.BarLen))
	}

	fmt.Fprintln(a.Out)
}

func (a *Agg) PushNum(ts string, value float64) {
	if a.Curr != ts {
		if a.Curr != "" {
			a.printNum()
		}

		a.Curr = ts
		a.NumList = []float64{}
	}

	a.NumList = append(a.NumList, value)
}
