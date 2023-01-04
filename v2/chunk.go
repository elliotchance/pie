package pie

// Chunk splits the input and returns multi slices whose length equals chunkLength,
// except for the last slice which may contain fewer elements.
//
// Examples:
//
//   Chunk([1, 2, 3], 4) => [ [1, 2, 3] ]
//   Chunk([1, 2, 3], 3) => [ [1, 2, 3] ]
//   Chunk([1, 2, 3], 2) => [ [1, 2], [3] ]
//   Chunk([1, 2, 3], 1) => [ [1], [2], [3] ]
//   Chunk([], 1)        => [ [] ]
//   Chunk([1, 2, 3], 0) => panic: chunkLength should be greater than 0
func Chunk[T any](ss []T, chunkLength int) [][]T {
	if chunkLength <= 0 {
		panic("chunkLength should be greater than 0")
	}

	result := make([][]T, 0)
	l := len(ss)
	if l == 0 {
		return result
	}

	var step = l / chunkLength
	if step == 0 {
		result = append(result, ss)
		return result
	}
	var remain = l % chunkLength
	for i := 0; i < step; i++ {
		result = append(result, ss[i*chunkLength:(i+1)*chunkLength])
	}
	if remain != 0 {
		result = append(result, ss[step*chunkLength:l])
	}

	return result
}
