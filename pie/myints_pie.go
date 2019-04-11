package pie

func (ss myInts) Average() float64 {
	if l := int(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

func (ss myInts) Sum() (sum int) {
	for _, s := range ss {
		sum += s
	}

	return
}
