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
	Out     io.Writer
	Curr    string
	Counter map[string]int
	NumList []float64
}

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

func sum(nums []float64) float64 {
	var total float64

	for _, n := range nums {
		total += n
	}

	return total
}

func avg(nums []float64) float64 {
	var total float64

	for _, n := range nums {
		total += n
	}

	return total / float64(len(nums))
}

func max(nums []float64) float64 {
	var maxVal float64

	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}

	return maxVal
}

func min(nums []float64) float64 {
	minVal := nums[0]

	for _, n := range nums {
		if n < minVal {
			minVal = n
		}
	}

	return minVal
}

func (a *Agg) printNum() {
	var aggregated float64

	switch a.Type {
	case Sum:
		aggregated = sum(a.NumList)
	case Avg:
		aggregated = avg(a.NumList)
	case Max:
		aggregated = max(a.NumList)
	case Min:
		aggregated = min(a.NumList)
	}

	fmt.Fprintf(a.Out, "%s\t%f\n", a.Curr, aggregated)
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
