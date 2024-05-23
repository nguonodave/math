package stats

func Variance(values []float64) float64 {
	avg := Average(values)

	// get the deviation (X-u), square it, and find the sum of the squared deviation
	// X is each data in the set, u is the mean
	sqrd_deviation_sum := 0.0
	for _, num := range values {
		deviation := num-avg
		sqrd_deviation_sum += deviation*deviation
	}

	// divide the sum of the squared deviation with the length of the dataset
	variance := sqrd_deviation_sum / float64(len(values))

	return variance
}
