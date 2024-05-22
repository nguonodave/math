package stats

import "testing"

func TestAverage(t *testing.T) {
	values := []float64{2.0, 4.0}
	output := Average(values)
	expected := 3
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}
}