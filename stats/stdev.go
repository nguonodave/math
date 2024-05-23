package stats

import "math"

func Stdev(values []float64) float64 {
	variance := Variance(values)
	stdev := math.Sqrt(variance)

	return float64(stdev)
}
