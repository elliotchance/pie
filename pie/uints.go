package pie

//go:generate pie Uints.*
type Uints []uint

//go:generate pie myUints.Sum.Average
type myUints []uint
