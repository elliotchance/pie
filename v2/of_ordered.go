package pie

import (
	"context"
	"math/rand"

	"golang.org/x/exp/constraints"
)

// OfOrdered encapsulates a slice to be used in multiple chained operations.
// OfOrdered requires that elements be numerical or a string for certain
// operations to be performed.
func OfOrdered[T constraints.Ordered](ss []T) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{ss}
}

// OfOrderedSlice provides the proxy methods that operate on slices. If the last
// method in the chain does not return a single value, you can access the Result
// to get final slice.
type OfOrderedSlice[T constraints.Ordered] struct {
	Result []T
}

func (o OfOrderedSlice[T]) All(fn func(value T) bool) bool {
	return All(o.Result, fn)
}

func (o OfOrderedSlice[T]) Any(fn func(value T) bool) bool {
	return Any(o.Result, fn)
}

func (o OfOrderedSlice[T]) AreSorted() bool {
	return AreSorted(o.Result)
}

func (o OfOrderedSlice[T]) AreUnique() bool {
	return AreUnique(o.Result)
}

func (o OfOrderedSlice[T]) Bottom(n int) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Bottom(o.Result, n)}
}

func (o OfOrderedSlice[T]) Contains(lookingFor T) bool {
	return Contains(o.Result, lookingFor)
}

func (o OfOrderedSlice[T]) Diff(against []T) ([]T, []T) {
	return Diff(o.Result, against)
}

func (o OfOrderedSlice[T]) DropTop(n int) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{DropTop(o.Result, n)}
}

func (o OfOrderedSlice[T]) DropWhile(f func(s T) bool) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{DropWhile(o.Result, f)}
}

func (o OfOrderedSlice[T]) Each(fn func(T)) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Each(o.Result, fn)}
}

func (o OfOrderedSlice[T]) Equals(rhs []T) bool {
	return Equals(o.Result, rhs)
}

func (o OfOrderedSlice[T]) Filter(condition func(T) bool) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Filter(o.Result, condition)}
}

func (o OfOrderedSlice[T]) FilterNot(condition func(T) bool) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{FilterNot(o.Result, condition)}
}

func (o OfOrderedSlice[T]) FindFirstUsing(fn func(value T) bool) int {
	return FindFirstUsing(o.Result, fn)
}

func (o OfOrderedSlice[T]) First() T {
	return First(o.Result)
}

func (o OfOrderedSlice[T]) FirstOr(defaultValue T) T {
	return FirstOr(o.Result, defaultValue)
}

func (o OfOrderedSlice[T]) Float64s() []float64 {
	return Float64s(o.Result)
}

func (o OfOrderedSlice[T]) Group() map[T]int {
	return Group(o.Result)
}

func (o OfOrderedSlice[T]) Insert(index int, values ...T) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Insert(o.Result, index, values...)}
}

func (o OfOrderedSlice[T]) Intersect(slices ...[]T) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Intersect(o.Result, slices...)}
}

func (o OfOrderedSlice[T]) Ints() []int {
	return Ints(o.Result)
}

func (o OfOrderedSlice[T]) Join(glue string) string {
	return Join(o.Result, glue)
}

func (o OfOrderedSlice[T]) JSONBytes() []byte {
	return JSONBytes(o.Result)
}

func (o OfOrderedSlice[T]) JSONBytesIndent(prefix, indent string) []byte {
	return JSONBytesIndent(o.Result, prefix, indent)
}

func (o OfOrderedSlice[T]) JSONString() string {
	return JSONString(o.Result)
}

func (o OfOrderedSlice[T]) JSONStringIndent(prefix, indent string) string {
	return JSONStringIndent(o.Result, prefix, indent)
}

func (o OfOrderedSlice[T]) Last() T {
	return Last(o.Result)
}

func (o OfOrderedSlice[T]) LastOr(defaultValue T) T {
	return LastOr(o.Result, defaultValue)
}

func (o OfOrderedSlice[T]) Map(fn func(T) T) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Map(o.Result, fn)}
}

func (o OfOrderedSlice[T]) Max() T {
	return Max(o.Result)
}

func (o OfOrderedSlice[T]) Min() T {
	return Min(o.Result)
}

func (o OfOrderedSlice[T]) Mode() OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Mode(o.Result)}
}

func (o OfOrderedSlice[T]) Reverse() OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Reverse(o.Result)}
}

func (o OfOrderedSlice[T]) Send(ctx context.Context, ch chan<- T) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Send(ctx, o.Result, ch)}
}

func (o OfOrderedSlice[T]) SequenceUsing(creator func(int) T, params ...int) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{SequenceUsing(o.Result, creator, params...)}
}

func (o OfOrderedSlice[T]) Shuffle(source rand.Source) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Shuffle(o.Result, source)}
}

func (o OfOrderedSlice[T]) Sort() OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Sort(o.Result)}
}

func (o OfOrderedSlice[T]) SortStableUsing(less func(a, b T) bool) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{SortStableUsing(o.Result, less)}
}

func (o OfOrderedSlice[T]) SortUsing(less func(a, b T) bool) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{SortUsing(o.Result, less)}
}

func (o OfOrderedSlice[T]) Strings() []string {
	return Strings(o.Result)
}

func (o OfOrderedSlice[T]) StringsUsing(transform func(T) string) []string {
	return StringsUsing(o.Result, transform)
}

func (o OfOrderedSlice[T]) SubSlice(start int, end int) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{SubSlice(o.Result, start, end)}
}

func (o OfOrderedSlice[T]) Top(n int) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Top(o.Result, n)}
}

func (o OfOrderedSlice[T]) Unique() OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Unique(o.Result)}
}

func (o OfOrderedSlice[T]) Unshift(elements ...T) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Unshift(o.Result, elements...)}
}

func (o OfOrderedSlice[T]) Delete(idx ...int) OfOrderedSlice[T] {
	return OfOrderedSlice[T]{Delete(o.Result, idx...)}
}
