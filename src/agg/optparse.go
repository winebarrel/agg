package agg

import (
	"flag"
	"fmt"
	"strings"
)

func ParseFlag(a *Agg) (err error) {
	var t string
	flag.StringVar(&t, "t", "count", "count,sum,avg,max,min")
	flag.Parse()

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
		return fmt.Errorf("Invalid aggregate type; %s", t)
	}

	return
}
