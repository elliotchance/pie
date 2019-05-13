package functions

import "fmt"

const (
	ForNumbers = 1 << iota
	ForStrings
	ForStructs
	ForMaps

	ForAll               = ForNumbers | ForStrings | ForStructs
	ForNumbersAndStrings = ForNumbers | ForStrings
)

// Function is a list of functions and which types they are available to. It is
// a slice instead of a may to make sure we iterate in a predicable order so
// regenerating files is deterministic.
var Functions = []struct {
	Name string
	File string
	For  int
}{
	{"Abs", "abs.go", ForNumbers},
	{"All", "all.go", ForAll},
	{"Any", "any.go", ForAll},
	{"Append", "append.go", ForAll},
	{"AreSorted", "are_sorted.go", ForNumbersAndStrings},
	{"AreUnique", "are_unique.go", ForNumbersAndStrings},
	{"Average", "average.go", ForNumbers},
	{"Bottom", "bottom.go", ForAll},
	{"Contains", "contains.go", ForAll},
	{"Diff", "diff.go", ForAll},
	{"Each", "each.go", ForAll},
	{"Extend", "extend.go", ForAll},
	{"Filter", "filter.go", ForAll},
	{"FilterNot", "filter_not.go", ForAll},
	{"FindFirst", "find_first.go", ForAll},
	{"First", "first.go", ForAll},
	{"FirstOr", "first_or.go", ForAll},
	{"Float64s", "float64s.go", ForAll},
	{"Intersect", "intersect.go", ForNumbersAndStrings},
	{"Ints", "ints.go", ForAll},
	{"Join", "join.go", ForStrings},
	{"JSONString", "json_string.go", ForAll},
	{"Keys", "keys.go", ForMaps},
	{"Last", "last.go", ForAll},
	{"LastOr", "last_or.go", ForAll},
	{"Len", "len.go", ForAll},
	{"Map", "map.go", ForAll},
	{"Max", "max.go", ForNumbersAndStrings},
	{"Median", "median.go", ForNumbers},
	{"Min", "min.go", ForNumbersAndStrings},
	{"Product", "product.go", ForNumbers},
	{"Random", "random.go", ForAll},
	{"Reduce", "reduce.go", ForNumbersAndStrings},
	{"Reverse", "reverse.go", ForAll},
	{"Send", "send.go", ForAll},
	{"Sequence", "sequence.go", ForNumbers},
	{"Shuffle", "shuffle.go", ForAll},
	{"Sort", "sort.go", ForNumbersAndStrings},
	{"SortStableUsing", "sort_stable_using.go", ForStrings | ForStructs},
	{"SortUsing", "sort_using.go", ForStrings | ForStructs},
	{"Strings", "strings.go", ForAll},
	{"Sum", "sum.go", ForNumbers},
	{"Top", "top.go", ForAll},
	{"ToStrings", "to_strings.go", ForAll},
	{"Unique", "unique.go", ForNumbersAndStrings},
	{"Values", "values.go", ForMaps},
}

type ElementType float64
type SliceType []ElementType
type StringElementType string
type StringSliceType []StringElementType
type KeyType string
type KeySliceType []KeyType
type MapType map[KeyType]ElementType

var ElementZeroValue ElementType

func (a ElementType) Equals(b ElementType) bool {
	return a == b
}

func (a ElementType) String() string {
	return fmt.Sprintf("%f", a)
}
