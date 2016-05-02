package collect

type Histogram map[int]int

func (s Series) Histogram() Histogram {
	var hist = Histogram{}
	for _, class := range s {
		hist[class.Value()] = hist[class.Value()] + 1
	}
	return hist
}

func (a Histogram) In(b Histogram) bool {
	for key, val := range a {
		if val != b[key] {
			return false
		}
	}
	return true
}
