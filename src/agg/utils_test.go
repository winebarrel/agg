package agg

import (
	. "."
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregatSum(t *testing.T) {
	assert := assert.New(t)
	expected := 6.0
	actual := AggregatSum([]float64{1, 2, 3})
	assert.Equal(expected, actual)
}

func TestAggregatAvg(t *testing.T) {
	assert := assert.New(t)
	expected := 2.0
	actual := AggregatAvg([]float64{1, 2, 3})
	assert.Equal(expected, actual)
}

func TestAggregatMax(t *testing.T) {
	assert := assert.New(t)
	expected := 3.0
	actual := AggregatMax([]float64{1, 2, 3})
	assert.Equal(expected, actual)
}

func TestAggregatMin(t *testing.T) {
	assert := assert.New(t)
	expected := 1.0
	actual := AggregatMin([]float64{1, 2, 3})
	assert.Equal(expected, actual)
}

func TestBar(t *testing.T) {
	assert := assert.New(t)
	width := 50
	max := 300
	assert.Equal("", Bar(0, width, max))
	assert.Equal("#", Bar(10, width, max))
	assert.Equal("#####", Bar(30, width, max))
	assert.Equal("############################", Bar(170, width, max))
	assert.Equal("##################################################", Bar(300, width, max))
}
