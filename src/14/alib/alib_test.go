package alib

import "testing"

var Debug bool = false

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("Skip: Debug is true")
	}
	v := Average([]float64{1.0, 2.0, 3.0, 4.0, 5.0})

	if v != 3.0 {
		t.Errorf("Average() = %v; want 3.0", v)
	}
}