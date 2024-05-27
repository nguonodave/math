package stats

import (
	"math"
	"testing"
)

func TestMedian(t *testing.T) {
	values := []float64{2.0, 4.0, 3.1, 2.3, 2.4, 1.3}
	output := math.Round(Median(values))
	
	expected := 2.0
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}
}
