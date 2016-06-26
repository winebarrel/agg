package agg

func AggregatSum(nums []float64) float64 {
	var total float64

	for _, n := range nums {
		total += n
	}

	return total
}

func AggregatAvg(nums []float64) float64 {
	var total float64

	for _, n := range nums {
		total += n
	}

	return total / float64(len(nums))
}

func AggregatMax(nums []float64) float64 {
	var maxVal float64

	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}

	return maxVal
}

func AggregatMin(nums []float64) float64 {
	minVal := nums[0]

	for _, n := range nums {
		if n < minVal {
			minVal = n
		}
	}

	return minVal
}

func Bar(cnt float64, width int, max int) string {
	if cnt < 1 {
		return ""
	}

	fraction := float64(width) / float64(max)

	if fraction > 1.0 {
		fraction = 1.0
	}

	length := int(float64(cnt) * fraction)
	bar := make([]byte, 0, length)

	for i := 0; i < length; i++ {
		bar = append(bar, '#')
	}

	return string(bar)
}
