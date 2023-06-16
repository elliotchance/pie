package pie

import (
	"context"
	"math/rand"

	"golang.org/x/exp/constraints"
)

// OfNumeric encapsulates a slice to be used in multiple chained operations.
// OfNumeric requires that elements be numerical for certain operations to be
// performed.
func OfNumeric[T constraints.Integer | constraints.Float](ss []T) OfNumericSlice[T] {
	return OfNumericSlice[T]{ss}
}

// OfNumericSlice provides the proxy methods that operate on slices. If the last
// method in the chain does not return a single value, you can access the Result
// to get final slice.
type OfNumericSlice[T constraints.Integer | constraints.Float] struct {
	Result []T
}

func (o OfNumericSlice[T]) All(fn func(value T) bool) bool {
	return All(o.Result, fn)
}

func (o OfNumericSlice[T]) Any(fn func(value T) bool) bool {
	return Any(o.Result, fn)
}

func (o OfNumericSlice[T]) AreSorted() bool {
	return AreSorted(o.Result)
}

func (o OfNumericSlice[T]) AreUnique() bool {
	return AreUnique(o.Result)
}

func (o OfNumericSlice[T]) Average() float64 {
	return Average(o.Result)
}

func (o OfNumericSlice[T]) Bottom(n int) OfNumericSlice[T] {
	return OfNumericSlice[T]{Bottom(o.Result, n)}
}

func (o OfNumericSlice[T]) Contains(lookingFor T) bool {
	return Contains(o.Result, lookingFor)
}

func (o OfNumericSlice[T]) Diff(against []T) ([]T, []T) {
	return Diff(o.Result, against)
}

func (o OfNumericSlice[T]) DropTop(n int) OfNumericSlice[T] {
	return OfNumericSlice[T]{DropTop(o.Result, n)}
}

func (o OfNumericSlice[T]) DropWhile(f func(s T) bool) OfNumericSlice[T] {
	return OfNumericSlice[T]{DropWhile(o.Result, f)}
}

func (o OfNumericSlice[T]) Each(fn func(T)) OfNumericSlice[T] {
	return OfNumericSlice[T]{Each(o.Result, fn)}
}

func (o OfNumericSlice[T]) Equals(rhs []T) bool {
	return Equals(o.Result, rhs)
}

func (o OfNumericSlice[T]) Filter(condition func(T) bool) OfNumericSlice[T] {
	return OfNumericSlice[T]{Filter(o.Result, condition)}
}

func (o OfNumericSlice[T]) FilterNot(condition func(T) bool) OfNumericSlice[T] {
	return OfNumericSlice[T]{FilterNot(o.Result, condition)}
}

func (o OfNumericSlice[T]) FindFirstUsing(fn func(value T) bool) int {
	return FindFirstUsing(o.Result, fn)
}

func (o OfNumericSlice[T]) First() T {
	return First(o.Result)
}

func (o OfNumericSlice[T]) FirstOr(defaultValue T) T {
	return FirstOr(o.Result, defaultValue)
}

func (o OfNumericSlice[T]) Float64s() []float64 {
	return Float64s(o.Result)
}

func (o OfNumericSlice[T]) Group() map[T]int {
	return Group(o.Result)
}

func (o OfNumericSlice[T]) Insert(index int, values ...T) OfNumericSlice[T] {
	return OfNumericSlice[T]{Insert(o.Result, index, values...)}
}

func (o OfNumericSlice[T]) Intersect(slices ...[]T) OfNumericSlice[T] {
	return OfNumericSlice[T]{Intersect(o.Result, slices...)}
}

func (o OfNumericSlice[T]) Ints() []int {
	return Ints(o.Result)
}

func (o OfNumericSlice[T]) Join(glue string) string {
	return Join(o.Result, glue)
}

func (o OfNumericSlice[T]) JSONBytes() []byte {
	return JSONBytes(o.Result)
}

func (o OfNumericSlice[T]) JSONBytesIndent(prefix, indent string) []byte {
	return JSONBytesIndent(o.Result, prefix, indent)
}

func (o OfNumericSlice[T]) JSONString() string {
	return JSONString(o.Result)
}

func (o OfNumericSlice[T]) JSONStringIndent(prefix, indent string) string {
	return JSONStringIndent(o.Result, prefix, indent)
}

func (o OfNumericSlice[T]) Last() T {
	return Last(o.Result)
}

func (o OfNumericSlice[T]) LastOr(defaultValue T) T {
	return LastOr(o.Result, defaultValue)
}

func (o OfNumericSlice[T]) Map(fn func(T) T) OfNumericSlice[T] {
	return OfNumericSlice[T]{Map(o.Result, fn)}
}

func (o OfNumericSlice[T]) Max() T {
	return Max(o.Result)
}

func (o OfNumericSlice[T]) Median() T {
	return Median(o.Result)
}

func (o OfNumericSlice[T]) Min() T {
	return Min(o.Result)
}

func (o OfNumericSlice[T]) Mode() OfNumericSlice[T] {
	return OfNumericSlice[T]{Mode(o.Result)}
}

func (o OfNumericSlice[T]) Product() T {
	return Product(o.Result)
}

func (o OfNumericSlice[T]) Random(source rand.Source) T {
	return Random(o.Result, source)
}

func (o OfNumericSlice[T]) Reduce(reducer func(T, T) T) T {
	return Reduce(o.Result, reducer)
}

func (o OfNumericSlice[T]) Reverse() OfNumericSlice[T] {
	return OfNumericSlice[T]{Reverse(o.Result)}
}

func (o OfNumericSlice[T]) Send(ctx context.Context, ch chan<- T) OfNumericSlice[T] {
	return OfNumericSlice[T]{Send(ctx, o.Result, ch)}
}

func (o OfNumericSlice[T]) Sequence(params ...int) OfNumericSlice[T] {
	return OfNumericSlice[T]{Sequence(o.Result, params...)}
}

func (o OfNumericSlice[T]) SequenceUsing(creator func(int) T, params ...int) OfNumericSlice[T] {
	return OfNumericSlice[T]{SequenceUsing(o.Result, creator, params...)}
}

func (o OfNumericSlice[T]) Shuffle(source rand.Source) OfNumericSlice[T] {
	return OfNumericSlice[T]{Shuffle(o.Result, source)}
}

func (o OfNumericSlice[T]) Sort() OfNumericSlice[T] {
	return OfNumericSlice[T]{Sort(o.Result)}
}

func (o OfNumericSlice[T]) SortStableUsing(less func(a, b T) bool) OfNumericSlice[T] {
	return OfNumericSlice[T]{SortStableUsing(o.Result, less)}
}

func (o OfNumericSlice[T]) SortUsing(less func(a, b T) bool) OfNumericSlice[T] {
	return OfNumericSlice[T]{SortUsing(o.Result, less)}
}

func (o OfNumericSlice[T]) Stddev() float64 {
	return Stddev(o.Result)
}

func (o OfNumericSlice[T]) Strings() []string {
	return Strings(o.Result)
}

func (o OfNumericSlice[T]) StringsUsing(transform func(T) string) []string {
	return StringsUsing(o.Result, transform)
}

func (o OfNumericSlice[T]) SubSlice(start int, end int) OfNumericSlice[T] {
	return OfNumericSlice[T]{SubSlice(o.Result, start, end)}
}

func (o OfNumericSlice[T]) Sum() T {
	return Sum(o.Result)
}

func (o OfNumericSlice[T]) Top(n int) OfNumericSlice[T] {
	return OfNumericSlice[T]{Top(o.Result, n)}
}

func (o OfNumericSlice[T]) Unique() OfNumericSlice[T] {
	return OfNumericSlice[T]{Unique(o.Result)}
}

func (o OfNumericSlice[T]) Unshift(elements ...T) OfNumericSlice[T] {
	return OfNumericSlice[T]{Unshift(o.Result, elements...)}
}

func (o OfNumericSlice[T]) Delete(idx ...int) OfNumericSlice[T] {
	return OfNumericSlice[T]{Delete(o.Result, idx...)}
}
