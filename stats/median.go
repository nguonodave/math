package stats

import (
	"math"
	"sort"
)

func Median(values []float64) int {

	var median int
	values_len := len(values)

	if len(values)%2 != 0 {
		sort.Float64s(values)
		median = int(math.Round(values[((values_len)/2)]))
	}
	return median
}
