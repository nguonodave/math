package stats

import "math"

func Average(values []float64) int {
	sum := 0.0
	for _, num := range values {
		sum += num
	}
	average := int(math.Round(sum / float64(len(values))))
	return average
}
