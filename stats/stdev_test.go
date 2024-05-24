package stats

import (
	"math"
	"testing"
)

func TestStdev(t *testing.T) {
	values := []float64{5.0, 6.0, 8.0, 9.0, 10.0, 11.0, 14.0}
	output := math.Round(Stdev(values))
	expected := 3.0
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}
}
