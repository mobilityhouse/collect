package collect

import ()

type Item int

/*
type Item interface {
	Value() int
}
*/
type Series []Item

func (s Series) Len() int           { return len(s) }
func (s Series) Less(i, j int) bool { return s[i] < s[j] }
func (s Series) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type IntItem Item
type IntItemSeries []IntItem

func (i Item) Value() int    { return int(i) }
func (i IntItem) Value() int { return int(i) }
