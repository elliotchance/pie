package functions

import "github.com/elliotchance/pie/pie/util"

// Sequence generates all numbers in range or returns nil if params invalid
//
// There are 3 variations to generate:
// 		1. [0, n).
//		2. [min, max).
//		3. [min, max) with step.
//
// if len(params) == 1 considered that will be returned slice between 0 and n,
// where n is the first param, [0, n).
// if len(params) == 2 considered that will be returned slice between min and max,
// where min is the first param, max is the second, [min, max).
// if len(params) > 2 considered that will be returned slice between min and max with step,
// where min is the first param, max is the second, step is the third one, [min, max) with step,
// others params will be ignored
func (ss SliceType) Sequence(params ...int) SliceType {
	var seq = func(min, max, step int) (seq SliceType) {
		lenght := int(util.Round(float64(max-min) / float64(step)))
		if lenght < 1 {
			return
		}

		seq = make(SliceType, lenght)
		for i := 0; i < lenght; min += step {
			seq[i] = ElementType(min)
			i++
		}

		return seq
	}

	if len(params) > 2 {
		return seq(params[0], params[1], params[2])
	} else if len(params) == 2 {
		return seq(params[0], params[1], 1)
	} else if len(params) == 1 {
		return seq(0, params[0], 1)
	} else {
		return ss
	}
}
