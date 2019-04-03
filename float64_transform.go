package pie

// Float64TransformFunc transforms an float64 value.
type Float64TransformFunc func(float64) float64

// AddFloat64 is addition.
func AddFloat64(a float64) Float64TransformFunc {
	return func(i float64) float64 {
		return a + i
	}
}
