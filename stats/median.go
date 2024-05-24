package stats

import "sort"

func Median(values []float64) float64 {
	var median float64
	values_len := len(values)
	sort.Float64s(values)

	if len(values)%2 != 0 {
		median = values[values_len/2]
	} else {
		median = Average(values[values_len/2-1 : values_len/2+1])
	}
	return median
}
