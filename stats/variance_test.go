package stats

import (
	"math"
	"testing"
)

func TestVariance(t *testing.T) {
	values := []float64{5.0, 6.0, 8.0, 9.0, 10.0, 11.0, 14.0}
	output := math.Round(Variance(values))
	
	expected := 8.0
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}
}
