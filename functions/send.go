package functions

import (
	"context"
)

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it closes channel after work
// returns quantity of sended elements if that != len(slice) considered func was canceled
func (ss SliceType) Send(ctx context.Context, ch chan<- ElementType) int {
	var amountSendedElements = 0

DONE:
	for _, s := range ss {
		select {
		case <-ctx.Done():
			break DONE
		default:
			ch <- s
			amountSendedElements++
		}
	}

	close(ch)
	return amountSendedElements
}
