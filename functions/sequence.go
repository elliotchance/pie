package functions

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
	var creator = func(i int) ElementType {
		return ElementType(i)
	}

	if len(params) > 2 {
		return ss.SequenceUsing(creator, params[0], params[1], params[2])
	} else if len(params) == 2 {
		return ss.SequenceUsing(creator, params[0], params[1])
	} else if len(params) == 1 {
		return ss.SequenceUsing(creator, params[0])
	} else {
		return nil
	}
}
