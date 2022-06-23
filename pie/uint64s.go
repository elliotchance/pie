package pie

//go:generate pie Uint64s.*
type Uint64s []uint64

//go:generate pie myUint64s.Sum.Average
type myUint64s []uint64
