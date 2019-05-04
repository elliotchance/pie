package pie

//go:generate pie Ints.*
type Ints []int

//go:generate pie myInts.Sum.Average
type myInts []int
