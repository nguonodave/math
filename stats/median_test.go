package stats

import "testing"

func TestMedian(t *testing.T) {
	values := []float64{2.0, 4.0, 3.1, 2.3, 2.4}
	output := Median(values)
	expected := 2
	if output != expected {
		t.Errorf("Expected %v got %v", expected, output)
	}
}