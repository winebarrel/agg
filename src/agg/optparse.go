package agg

import (
	"flag"
	"fmt"
	"strings"
)

func ParseFlag(a *Agg) (err error) {
	var t string
	flag.StringVar(&t, "t", "count", "count,sum,avg,max,min")
	flag.IntVar(&a.TrimLen, "s", 0, "Trim timestamp suffix length")
	flag.IntVar(&a.BarLen, "bar", 0, "ASCII bar length")
	flag.Parse()

	if a.TrimLen < 0 {
		return fmt.Errorf("Invalid trim length: %d", a.TrimLen)
	}

	if a.BarLen < 0 {
		return fmt.Errorf("Invalid bar length: %d", a.BarLen)
	}

	switch strings.Title(t) {
	case Count.String():
		a.Type = Count
	case Sum.String():
		a.Type = Sum
	case Avg.String():
		a.Type = Avg
	case Max.String():
		a.Type = Max
	case Min.String():
		a.Type = Min
	default:
		return fmt.Errorf("Invalid aggregate type: %s", t)
	}

	return
}
