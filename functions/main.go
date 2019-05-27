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
	BigO string
}{
	{"Abs", "abs.go", ForNumbers, "n"},
	{"All", "all.go", ForAll, "n"},
	{"Any", "any.go", ForAll, "n"},
	{"Append", "append.go", ForAll, "n"},
	{"AreSorted", "are_sorted.go", ForNumbersAndStrings, "n"},
	{"AreUnique", "are_unique.go", ForNumbersAndStrings, "n"},
	{"Average", "average.go", ForNumbers, "n"},
	{"Bottom", "bottom.go", ForAll, "n"},
	{"Contains", "contains.go", ForAll, "n"},
	{"Diff", "diff.go", ForAll, "n²"},
	{"DropTop", "drop_top.go", ForAll, "n"},
	{"Each", "each.go", ForAll, "n"},
	{"Equals", "equals.go", ForAll, "n"},
	{"Extend", "extend.go", ForAll, "n"},
	{"Filter", "filter.go", ForAll, "n"},
	{"FilterNot", "filter_not.go", ForAll, "n"},
	{"FindFirstUsing", "find_first_using.go", ForAll, "n"},
	{"First", "first.go", ForAll, "1"},
	{"FirstOr", "first_or.go", ForAll, "1"},
	{"Float64s", "float64s.go", ForAll, "n"},
	{"Intersect", "intersect.go", ForNumbersAndStrings, "n"},
	{"Ints", "ints.go", ForAll, "n"},
	{"Join", "join.go", ForStrings, "n"},
	{"JSONBytes", "json_bytes.go", ForAll, "n"},
	{"JSONBytesIndent", "json_bytes_indent.go", ForAll, "n"},
	{"JSONString", "json_string.go", ForAll, "n"},
	{"JSONStringIndent", "json_string_indent.go", ForAll, "n"},
	{"Keys", "keys.go", ForMaps, "n"},
	{"Last", "last.go", ForAll, "1"},
	{"LastOr", "last_or.go", ForAll, "1"},
	{"Len", "len.go", ForAll, "1"},
	{"Map", "map.go", ForAll, "n"},
	{"Max", "max.go", ForNumbersAndStrings, "n"},
	{"Median", "median.go", ForNumbers, "n"},
	{"Min", "min.go", ForNumbersAndStrings, "n"},
	{"Mode", "mode.go", ForAll, "n"},
	{"Product", "product.go", ForNumbers, "n"},
	{"Random", "random.go", ForAll, "1"},
	{"Reduce", "reduce.go", ForNumbersAndStrings, "n"},
	{"Reverse", "reverse.go", ForAll, "n"},
	{"Send", "send.go", ForAll, "n"},
	{"Sequence", "sequence.go", ForNumbers, "n"},
	{"SequenceUsing", "sequence_using.go", ForAll, "n"},
	{"Shift", "shift.go", ForAll, "n"},
	{"Shuffle", "shuffle.go", ForAll, "n"},
	{"Sort", "sort.go", ForNumbersAndStrings, "n⋅log(n)"},
	{"SortStableUsing", "sort_stable_using.go", ForStrings | ForStructs, "n⋅log(n)"},
	{"SortUsing", "sort_using.go", ForStrings | ForStructs, "n⋅log(n)"},
	{"Strings", "strings.go", ForAll, "n"},
	{"SubSlice", "sub_slice.go", ForAll, "n"},
	{"Sum", "sum.go", ForNumbers, "n"},
	{"Top", "top.go", ForAll, "n"},
	{"StringsUsing", "strings_using.go", ForAll, "n"},
	{"Unique", "unique.go", ForNumbersAndStrings, "n"},
	{"Unshift", "unshift.go", ForAll, "n"},
	{"Values", "values.go", ForMaps, "n"},
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
