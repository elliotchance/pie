package pie

import (
	"golang.org/x/exp/constraints"
	"math"
)

// Stddev is the standard deviation
func Stddev[T constraints.Integer | constraints.Float](ss []T) float64 {
	if len(ss) == 0 {
		return 0.0
	}

	avg := Average(ss)

	var sd float64
	for i := range ss {
		sd += math.Pow(float64(ss[i])-avg, 2)
	}
	sd = math.Sqrt(sd / float64(len(ss)))

	return sd
}
