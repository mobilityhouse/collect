package collect

import (
	"testing"
)

func TestHistogramBagsItems(t *testing.T) {
	var series = Series([]Item{1, 1, 1, 2, 2, 3, 5})
	var histogram = series.Histogram()
	if histogram[0] != 0 {
		t.Errorf("Histogram should not include 0")
	}
	if histogram[1] != 3 {
		t.Errorf("Histogram should include three 1's")
	}
	if histogram[2] != 2 {
		t.Errorf("Histogram should include three 1's")
	}
	if histogram[3] != 1 {
		t.Errorf("Histogram should include one 3")
	}
	if histogram[4] != 0 {
		t.Errorf("Histogram should not include 4")
	}
	if histogram[5] != 1 {
		t.Errorf("Histogram should include one 5")
	}
}

func TestHistogramIn(t *testing.T) {
	a, b := Series([]Item{0, 1, 1, 3}), Series([]Item{0, 1, 1, 3})
	if !a.Histogram().In(b.Histogram()) {
		t.Errorf("Histograms should have been equal")
	}

	c := Series([]Item{0, 1, 1, 3, 4})
	if !a.Histogram().In(c.Histogram()) {
		t.Errorf("Histogram a should be contained in histogram b")
	}

	d, e := Series([]Item{1, 1, 3}), Series([]Item{})
	if a.Histogram().In(d.Histogram()) {
		t.Errorf("Histogram a is not completely contained in histogram d")
	}

	if a.Histogram().In(e.Histogram()) {
		t.Errorf("Histogram a is not completely contained in histogram d")
	}
}
