package pie

import (
	"context"
)

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it doesn't close channel after work
// returns sent elements if len(this) != len(old) considered func was canceled
func Send[T any](ctx context.Context, ss []T, ch chan<- T) []T {
	for i, s := range ss {
		select {
		case <-ctx.Done():
			return ss[:i]
		default:
			ch <- s
		}
	}

	return ss
}
