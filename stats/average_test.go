package stats

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	values := []float64{2.0, 4.0}
	output := math.Round(Average(values))
	expected := 3.0
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}
}