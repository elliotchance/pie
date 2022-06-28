package pie

import "fmt"

// Chunk Split slice to chunks
func Chunk[T any](ss []T, chunkSize int, callback func(ss []T) (stopped bool)) {
	if chunkSize <= 0 {
		panic(fmt.Sprintf("invalid chunk size %d", chunkSize))
	}
	var stopped bool
	for i := 0; i < len(ss); i += chunkSize {
		if i+chunkSize < len(ss) {
			stopped = callback(ss[i : i+chunkSize])
		} else {
			stopped = callback(ss[i:])
		}
		if stopped {
			break
		}
	}
}
