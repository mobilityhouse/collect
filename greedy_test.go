package collect

import (
	"testing"
)

// Optimize for requested values in order to minimize for shortages in the
// system. Every system should get as close to the available value requested.
func AssertGreedyCollectionContent(t *testing.T, requested Item, available, expected Series) {
	var collected Series
	collected = available.GreedyCollect(Item(requested))
	if !collected.Histogram().In(expected.Histogram()) {
		t.Errorf("From %v, expected %v but received %v", available, expected, collected)
	}
}

func TestOptimizeForRequest(t *testing.T) {
	requested := Item(12)
	AssertGreedyCollectionContent(t, requested, []Item{100, 11}, []Item{11})
	AssertGreedyCollectionContent(t, requested, []Item{13, 20}, []Item{13})
	AssertGreedyCollectionContent(t, requested, []Item{10, 6, 6}, []Item{6, 6})
}

// Favor offering from multiple sources for improved reliability (redundancy)
func TestOptimizeForSpread(t *testing.T) {
	requested := Item(12)
	AssertGreedyCollectionContent(t, requested, []Item{12, 6, 6}, []Item{6, 6})
	AssertGreedyCollectionContent(t, requested, []Item{11, 4, 4, 4, 1}, []Item{4, 4, 4})
	AssertGreedyCollectionContent(t, requested, []Item{6, 4, 2, 2, 1}, []Item{6, 4, 2})
	AssertGreedyCollectionContent(t, requested, []Item{6, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []Item{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	AssertGreedyCollectionContent(t, requested, []Item{12, 12, 12, 12, 12, 12, 12, 12, 6, 3, 1, 2}, []Item{6, 3, 1, 2})

	// Both 10, 1, 1 and 1, 9, 2 offer a solution distributed over 3 sources,
	// however; the first solution relies too heavily on a single unit which
	// represents a greater risk on failure.
	AssertGreedyCollectionContent(t, requested, []Item{10, 1, 1, 9, 2}, []Item{1, 9, 2})
	AssertGreedyCollectionContent(t, requested, []Item{7, 3, 2, 5, 2}, []Item{3, 2, 5, 2})

	AssertGreedyCollectionContent(t, requested, []Item{10, 1, 1, 5, 6}, []Item{1, 5, 6})
}

/*
func TestTypes(t *testing.T) {
	Report([]Item{})
}
*/
