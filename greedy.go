// Package collect offers judging logic to aid in selecting offers
package collect

import (
	"fmt"
	"sort"
)

// GreedyCollect greedily collects from Series to approximate requested Item.
// The greedy collection strategy basically just iterates through the series
// and accumulates encountered items until the requested value is met. Naivety
// is tackled by sorting the items fist. This ensures that at any point in the
// series, if the collection of the current item were to exceed the requested
// value, the same result would hold for any subsequent items from that point
// onward (since every item from that point is at least as small as the current
// item).
func (available Series) GreedyCollect(requested Item) Series {
	var series = Series(available)
	sort.Sort(series)
	fmt.Printf("series: %v\n", series)

	acc := int(0)
	for key, item := range series {
		if requested < Item(acc+item.Value()) {
			return series[0:key]
		}
		acc = acc + item.Value()
	}

	return Series([]Item{})
}

func CollectByHistogram(requested Item, available Histogram) Histogram {
	return Histogram{}
}
