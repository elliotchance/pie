package pie

import (
	"context"
	"math/rand"
)

// Of encapsulates a slice to be used in multiple chained operations.
func Of[T any](ss []T) OfSlice[T] {
	return OfSlice[T]{ss}
}

// OfSlice provides the proxy methods that operate on slices. If the last method
// in the chain does not return a single value, you can access the Result to get
// final slice.
type OfSlice[T any] struct {
	Result []T
}

func (o OfSlice[T]) All(fn func(value T) bool) bool {
	return All(o.Result, fn)
}

func (o OfSlice[T]) Any(fn func(value T) bool) bool {
	return Any(o.Result, fn)
}

func (o OfSlice[T]) Bottom(n int) OfSlice[T] {
	return OfSlice[T]{Bottom(o.Result, n)}
}

func (o OfSlice[T]) DropTop(n int) OfSlice[T] {
	return OfSlice[T]{DropTop(o.Result, n)}
}

func (o OfSlice[T]) Each(fn func(T)) OfSlice[T] {
	return OfSlice[T]{Each(o.Result, fn)}
}

func (o OfSlice[T]) Filter(condition func(T) bool) OfSlice[T] {
	return OfSlice[T]{Filter(o.Result, condition)}
}

func (o OfSlice[T]) FilterNot(condition func(T) bool) OfSlice[T] {
	return OfSlice[T]{FilterNot(o.Result, condition)}
}

func (o OfSlice[T]) FindFirstUsing(fn func(value T) bool) int {
	return FindFirstUsing(o.Result, fn)
}

func (o OfSlice[T]) FirstOr(defaultValue T) T {
	return FirstOr(o.Result, defaultValue)
}

func (o OfSlice[T]) Insert(index int, values ...T) OfSlice[T] {
	return OfSlice[T]{Insert(o.Result, index, values...)}
}

func (o OfSlice[T]) LastOr(defaultValue T) T {
	return LastOr(o.Result, defaultValue)
}

func (o OfSlice[T]) Map(fn func(T) T) OfSlice[T] {
	return OfSlice[T]{Map(o.Result, fn)}
}

func (o OfSlice[T]) Reverse() OfSlice[T] {
	return OfSlice[T]{Reverse(o.Result)}
}

func (o OfSlice[T]) Send(ctx context.Context, ch chan<- T) OfSlice[T] {
	return OfSlice[T]{Send(ctx, o.Result, ch)}
}

func (o OfSlice[T]) SequenceUsing(creator func(int) T, params ...int) OfSlice[T] {
	return OfSlice[T]{SequenceUsing(o.Result, creator, params...)}
}

func (o OfSlice[T]) Shuffle(source rand.Source) OfSlice[T] {
	return OfSlice[T]{Shuffle(o.Result, source)}
}

func (o OfSlice[T]) SortUsing(less func(a, b T) bool) OfSlice[T] {
	return OfSlice[T]{SortUsing(o.Result, less)}
}

func (o OfSlice[T]) StringsUsing(transform func(T) string) []string {
	return StringsUsing(o.Result, transform)
}

func (o OfSlice[T]) SubSlice(start int, end int) OfSlice[T] {
	return OfSlice[T]{SubSlice(o.Result, start, end)}
}

func (o OfSlice[T]) Top(n int) OfSlice[T] {
	return OfSlice[T]{Top(o.Result, n)}
}

func (o OfSlice[T]) Unshift(elements ...T) OfSlice[T] {
	return OfSlice[T]{Unshift(o.Result, elements...)}
}
