package functions

const (
	ForNumbers = 1 << iota
	ForStrings
	ForStructs

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
	{"Append", "append.go", ForAll},
	{"AreSorted", "are_sorted.go", ForNumbersAndStrings},
	{"AreUnique", "are_unique.go", ForNumbersAndStrings},
	{"Average", "average.go", ForNumbers},
	{"Contains", "contains.go", ForAll},
	{"Extend", "extend.go", ForAll},
	{"First", "first.go", ForAll},
	{"FirstOr", "first_or.go", ForAll},
	{"Join", "join.go", ForStrings},
	{"JSONString", "json_string.go", ForAll},
	{"Last", "last.go", ForAll},
	{"LastOr", "last_or.go", ForAll},
	{"Len", "len.go", ForAll},
	{"Max", "max.go", ForNumbersAndStrings},
	{"Min", "min.go", ForNumbersAndStrings},
	{"Reverse", "reverse.go", ForAll},
	{"Select", "select.go", ForAll},
	{"Sort", "sort.go", ForNumbersAndStrings},
	{"Sum", "sum.go", ForNumbers},
	{"ToStrings", "to_strings.go", ForAll},
	{"Transform", "transform.go", ForAll},
	{"Unique", "unique.go", ForNumbersAndStrings},
	{"Unselect", "unselect.go", ForAll},
}

type ElementType float64
type SliceType []ElementType
type StringElementType string
type StringSliceType []StringElementType

var ElementZeroValue ElementType
