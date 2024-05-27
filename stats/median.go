package stats

import "sort"

func Median(values []float64) float64 {
	// initialize median variable to hold the median result
	// get the length of the data set
	// sort the values in the data set
	var median float64
	values_len := len(values)
	sort.Float64s(values)

	// if the length of the number of values is odd, simply get the value in the middle
	// if it's even, get the two values in the middle and find their mean
	if len(values)%2 != 0 {
		median = values[values_len/2]
	} else if len(values) > 0 && len(values)%2 == 0 {
		median = Average(values[values_len/2-1 : values_len/2+1])
	}

	return median
}
