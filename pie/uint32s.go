package pie

//go:generate pie Uint32s.*
type Uint32s []uint32

//go:generate pie myUint32s.Sum.Average
type myUint32s []uint32
