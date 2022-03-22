package pie_test

import (
	"context"
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var sendTests = []struct {
	ss            []float64
	recieveDelay  time.Duration
	canceledDelay time.Duration
	expected      []float64
}{
	{
		nil,
		0,
		0,
		nil,
	},
	{
		[]float64{1.2, 3.2},
		0,
		0,
		[]float64{1.2, 3.2},
	},
	{
		[]float64{1.2, 3.2},
		time.Millisecond * 30,
		time.Millisecond * 10,
		[]float64{1.2},
	},
	{
		[]float64{1.2, 3.2},
		time.Millisecond * 3,
		time.Millisecond * 10,
		[]float64{1.2, 3.2},
	},
}

func TestSend(t *testing.T) {
	for _, test := range sendTests {
		t.Run("", func(t *testing.T) {
			ch := make(chan float64)

			actual := getFloat64sFromChan(ch, test.recieveDelay)
			ctx := createContextByDelay(test.canceledDelay)

			actualSent := pie.Send(ctx, test.ss, ch)
			close(ch)

			assert.Equal(t, test.expected, actualSent)
			assert.Equal(t, test.expected, actual())
		})
	}
}

func getFloat64sFromChan(ch chan float64, t time.Duration) func() []float64 {
	done := make(chan struct{})
	var c []float64
	if t > 0 {
		go func() {
			ticker := time.NewTicker(t)
			defer ticker.Stop()
			for range ticker.C {
				val, ok := <-ch
				if !ok {
					break
				} else {
					c = append(c, val)
				}
			}
			done <- struct{}{}

		}()
	} else {
		go func() {
			for val := range ch {
				c = append(c, val)
			}
			done <- struct{}{}
		}()
	}

	return func() []float64 {
		<-done
		return c
	}
}

func createContextByDelay(t time.Duration) context.Context {
	ctx := context.Background()
	if t > 0 {
		ctx, _ = context.WithTimeout(ctx, t)
	}

	return ctx
}
