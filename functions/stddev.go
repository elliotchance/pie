package functions

import "math"

// Stddev is the standard deviation
func (ss SliceType) Stddev() float64 {
	if len(ss) == 0 {
		return 0.0
	}

	avg := ss.Average()

	var sd float64
	for i := range ss {
		sd += math.Pow(float64(ss[i])-avg, 2)
	}
	sd = math.Sqrt(sd / float64(len(ss)))

	return sd
}
