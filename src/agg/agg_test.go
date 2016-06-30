package agg

import (
	. "."
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	a := &Agg{Type: Count, Out: buf}

	a.PushStr("12:10", "FOO")
	a.PushStr("12:10", "FOO")
	a.PushStr("12:10", "BAR")
	a.PushStr("12:11", "FOO")
	a.PushStr("12:11", "BAR")
	a.PushStr("12:11", "BAR")
	a.PushStr("12:12", "ZOO")

	expected := `12:10	BAR:1	FOO:2
12:11	BAR:2	FOO:1
`
	actual := buf.String()

	assert.Equal(expected, actual)
}

func TestCountall(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	a := &Agg{Type: Countall, Out: buf}

	a.PushStr("12:10", "FOO")
	a.PushStr("12:10", "FOO")
	a.PushStr("12:10", "BAR")
	a.PushStr("12:11", "FOO")
	a.PushStr("12:11", "BAR")
	a.PushStr("12:11", "BAR")
	a.PushStr("12:12", "ZOO")

	expected := `12:10	3
12:11	3
`
	actual := buf.String()

	assert.Equal(expected, actual)
}

func TestSum(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	a := &Agg{Type: Sum, Out: buf}

	a.PushNum("12:10", 1)
	a.PushNum("12:10", 2)
	a.PushNum("12:10", 3)
	a.PushNum("12:11", 2)
	a.PushNum("12:11", 3)
	a.PushNum("12:11", 4)
	a.PushNum("12:12", 5)

	expected := `12:10	6.000000
12:11	9.000000
`
	actual := buf.String()

	assert.Equal(expected, actual)
}

func TestSumWithBar(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	a := &Agg{Type: Sum, Out: buf, BarLen: 10}

	a.PushNum("12:10", 1)
	a.PushNum("12:10", 2)
	a.PushNum("12:10", 3)
	a.PushNum("12:11", 2)
	a.PushNum("12:11", 3)
	a.PushNum("12:11", 4)
	a.PushNum("12:12", 5)

	expected := `12:10	6.000000	######
12:11	9.000000	#########
`
	actual := buf.String()

	assert.Equal(expected, actual)
}

func TestAvg(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	a := &Agg{Type: Avg, Out: buf}

	a.PushNum("12:10", 1)
	a.PushNum("12:10", 2)
	a.PushNum("12:10", 3)
	a.PushNum("12:11", 2)
	a.PushNum("12:11", 3)
	a.PushNum("12:11", 4)
	a.PushNum("12:12", 5)

	expected := `12:10	2.000000
12:11	3.000000
`
	actual := buf.String()

	assert.Equal(expected, actual)
}

func TestMax(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	a := &Agg{Type: Max, Out: buf}

	a.PushNum("12:10", 1)
	a.PushNum("12:10", 2)
	a.PushNum("12:10", 3)
	a.PushNum("12:11", 2)
	a.PushNum("12:11", 3)
	a.PushNum("12:11", 4)
	a.PushNum("12:12", 5)

	expected := `12:10	3.000000
12:11	4.000000
`
	actual := buf.String()

	assert.Equal(expected, actual)
}

func TestMin(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	a := &Agg{Type: Min, Out: buf}

	a.PushNum("12:10", 1)
	a.PushNum("12:10", 2)
	a.PushNum("12:10", 3)
	a.PushNum("12:11", 2)
	a.PushNum("12:11", 3)
	a.PushNum("12:11", 4)
	a.PushNum("12:12", 5)

	expected := `12:10	1.000000
12:11	2.000000
`
	actual := buf.String()

	assert.Equal(expected, actual)
}
