package pie

//go:generate pie Int64s.*
type Int64s []int

//go:generate pie myInt64s.Sum.Average
type myInt64s []int