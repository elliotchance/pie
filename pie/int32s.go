package pie

//go:generate pie Int32s.*
type Int32s []int32

//go:generate pie myInt32s.Sum.Average
type myInt32s []int32
