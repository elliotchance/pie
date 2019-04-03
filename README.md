# 🍕 `github.com/elliotchance/pie` [![GoDoc](https://godoc.org/github.com/elliotchance/pie?status.svg)](https://godoc.org/github.com/elliotchance/pie)

**Enjoy a slice!** `pie` is a utility library for dealing with slices that
focuses on type safety and performance.

It can be used with the Go-style package functions:

```go
names := []string{"Bob", "Sally", "John", "Jane"}
shortNames := pie.StringsOnly(names, func(s string) bool {
	return len(s) <= 3
})

// pie.Strings{"Bob"}
```

Or, they can be chained for more complex operations:

```go
pie.Strings{"Bob", "Sally", "John", "Jane"}.
	Without(pie.Prefix("J")).
	Transform(pie.ToUpper()).
	Last()

// "SALLY"
```

# Functions

## Slices

| Strings | Ints | Float64s | Description |     |
| ------- | ---- | -------- | ----------- | --- |
| | [`IntsAverage`](https://godoc.org/github.com/elliotchance/pie#IntsAverage) | [`Float64sAverage`](https://godoc.org/github.com/elliotchance/pie#Float64sAverage) | The average (mean) value, or a zeroed value. | O(n) |
| [`StringsContains`](https://godoc.org/github.com/elliotchance/pie#StringsContains) | [`IntsContains`](https://godoc.org/github.com/elliotchance/pie#IntsContains) | [`Float64sContains`](https://godoc.org/github.com/elliotchance/pie#Float64sContains) | Check if the value exists in the slice. | O(n) |
| [`StringsFirst`](https://godoc.org/github.com/elliotchance/pie#StringsFirst) | [`IntsFirst`](https://godoc.org/github.com/elliotchance/pie#IntsFirst) | [`Float64sFirst`](https://godoc.org/github.com/elliotchance/pie#Float64sFirst) | The first element, or a zeroed value. | O(1) |
| [`StringsFirstOr`](https://godoc.org/github.com/elliotchance/pie#StringsFirstOr) | [`IntsFirstOr`](https://godoc.org/github.com/elliotchance/pie#IntsFirstOr) | [`Float64sFirstOr`](https://godoc.org/github.com/elliotchance/pie#Float64sFirstOr) | The first element, or a default value. | O(1) |
| [`StringsLast`](https://godoc.org/github.com/elliotchance/pie#StringsLast) | [`IntsLast`](https://godoc.org/github.com/elliotchance/pie#IntsLast) | [`Float64sLast`](https://godoc.org/github.com/elliotchance/pie#Float64sLast) | The last element, or a zeroed value. | O(1) |
| [`StringsLastOr`](https://godoc.org/github.com/elliotchance/pie#StringsLastOr) | [`IntsLastOr`](https://godoc.org/github.com/elliotchance/pie#IntsLastOr) | [`Float64sLastOr`](https://godoc.org/github.com/elliotchance/pie#Float64sLastOr) | The last element, or a default value. | O(1) |
| [`StringsMax`](https://godoc.org/github.com/elliotchance/pie#StringsMax) | [`IntsMax`](https://godoc.org/github.com/elliotchance/pie#IntsMax) | [`Float64sMax`](https://godoc.org/github.com/elliotchance/pie#Float64sMax) | The maximum value, or a zeroes value. | O(n) |
| [`StringsMin`](https://godoc.org/github.com/elliotchance/pie#StringsMin) | [`IntsMin`](https://godoc.org/github.com/elliotchance/pie#IntsMin) | [`Float64sMin`](https://godoc.org/github.com/elliotchance/pie#Float64sMin) | The minimum value, or a zeroed value. | O(n) |
| [`StringsOnly`](https://godoc.org/github.com/elliotchance/pie#StringsOnly) | [`IntsOnly`](https://godoc.org/github.com/elliotchance/pie#IntsOnly) | [`Float64sOnly`](https://godoc.org/github.com/elliotchance/pie#Float64sOnly) | A new slice containing only the elements that returned true from the condition. | O(n) |
| | [`IntsSum`](https://godoc.org/github.com/elliotchance/pie#IntsSum) | [`Float64sSum`](https://godoc.org/github.com/elliotchance/pie#Float64sSum) | Sum (total) of all elements. | O(n) |
| [`StringsTransform`](https://godoc.org/github.com/elliotchance/pie#StringsTransform) | [`IntsTransform`](https://godoc.org/github.com/elliotchance/pie#IntsTransform) | [`Float64sTransform`](https://godoc.org/github.com/elliotchance/pie#Float64sTransform) | A new slice where each element has been transformed. | O(n) |
| [`StringsWithout`](https://godoc.org/github.com/elliotchance/pie#StringsWithout) | [`IntsWithout`](https://godoc.org/github.com/elliotchance/pie#IntsWithout) | [`Float64sWithout`](https://godoc.org/github.com/elliotchance/pie#Float64sWithout) | A new slice containing only the elements that returned false from the condition. | O(n) |

## Conditional

| Strings | Ints | Float64s | Description |
| ------- | ---- | -------- | ----------- |
| [`EqualString`](https://godoc.org/github.com/elliotchance/pie#EqualString) | [`EqualInt`](https://godoc.org/github.com/elliotchance/pie#EqualInt) | [`EqualFloat64`](https://godoc.org/github.com/elliotchance/pie#EqualFloat64) | Check if the values are equal. |
| [`GreaterThanString`](https://godoc.org/github.com/elliotchance/pie#GreaterThanString) | [`GreaterThanInt`](https://godoc.org/github.com/elliotchance/pie#EqualInt) | [`GreaterThanFloat64`](https://godoc.org/github.com/elliotchance/pie#GreaterThanFloat64) | Check if the values are greater than. |
| [`GreaterThanEqualString`](https://godoc.org/github.com/elliotchance/pie#GreaterThanEqualString) | [`GreaterThanEqualInt`](https://godoc.org/github.com/elliotchance/pie#GreaterThanEqualInt) | [`GreaterThanEqualFloat64`](https://godoc.org/github.com/elliotchance/pie#GreaterThanEqualFloat64) | Check if the values are greater than or equal to. |
| [`LessThanString`](https://godoc.org/github.com/elliotchance/pie#LessThanString) | [`LessThanInt`](https://godoc.org/github.com/elliotchance/pie#LessThanInt) | [`LessThanFloat64`](https://godoc.org/github.com/elliotchance/pie#LessThanFloat64) | Check if the values are less than. |
| [`LessThanEqualString`](https://godoc.org/github.com/elliotchance/pie#LessThanEqualString) | [`LessThanEqualInt`](https://godoc.org/github.com/elliotchance/pie#LessThanEqualInt) | [`LessThanEqualFloat64`](https://godoc.org/github.com/elliotchance/pie#LessThanEqualFloat64) | Check if the values are less than or equal to. |
| [`NotEqualString`](https://godoc.org/github.com/elliotchance/pie#NotEqualString) | [`NotEqualInt`](https://godoc.org/github.com/elliotchance/pie#NotEqualInt) | [`NotEqualFloat64`](https://godoc.org/github.com/elliotchance/pie#NotEqualFloat64) | Check if the values are not equal. |
| [`Prefix`](https://godoc.org/github.com/elliotchance/pie#Prefix) | | | Check if the string starts with another string. |
| [`Suffix`](https://godoc.org/github.com/elliotchance/pie#Suffix) | | | Check if the string ends with another string. |

## Transforms

Some of the functions listed here are part of the Go native libraries, but they
are useful to list.

### Strings

| Strings | Description |
| ------- | ----------- |
| [`strings.ToLower`](https://golang.org/pkg/strings/#ToLower) | Convert string to lower case. |
| [`strings.ToTitle`](https://golang.org/pkg/strings/#ToTitle) | Convert string to title case. |
| [`strings.ToUpper`](https://golang.org/pkg/strings/#ToUpper) | Convert string to upper case. |
| [`strings.TrimSpace`](https://golang.org/pkg/strings/#TrimSpace) | Trim leading and trailing whitespace. |

### Ints and Float64s

| Ints | Float64s | Description |
| ---- | -------- | ----------- |
| [`AddInt`](https://godoc.org/github.com/elliotchance/pie#AddInt) | [`AddFloat64`](https://godoc.org/github.com/elliotchance/pie#AddFloat64) | Addition. |

# FAQ

## How do I use it?

You can include it like any other package through your favourite package
manager:

1. Go modules ([you should be using this one](http://elliot.land/post/migrating-projects-from-dep-to-go-modules)):
`go get -u github.com/elliotchance/pie`

2. Dep: `dep ensure -add github.com/elliotchance/pie`

## Why do we need another library for this?

Yes, there are some other great options like
[`thoas/go-funk`](https://github.com/thoas/go-funk),
[`leaanthony/slicer`](https://github.com/leaanthony/slicer),
[`viant/toolbox`](https://github.com/viant/toolbox) and
[`alxrm/ugo`](https://github.com/alxrm/ugo) to name a few.

A lot of my work is dealing with servers that need to be high performance. I
found myself creating all the same utility functions like StringSliceContains
because I wanted to avoid reflection.

## What are the goals of `pie`?

1. **Type safety.** I never want to hit runtime bugs because I could pass in the
wrong type, or perform an invalid type case out the other end.

2. **Performance.** The functions need to be as fast as native Go
implementations otherwise there's no point in this library existing.

3. **Nil-safe.** All of the functions will happily accept nil and treat them as
empty slices. Apart from less possible panics, it makes it easier to chain.

There are some downsides with this approach:

1. It won't support all slice types. Sorry, you can use these actions on
`[]Foo`.

2. Until
[parametric polymorphism (generics) possibly arrives in Go v2](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)
there will need to be duplicate code in `pie` to compensate.

## Can I contribute?

Absolutely. Pull requests are always welcome. Your PR must include:

1. The package functions and chainable functions. Such as `pie.StringsSort` and `Strings.Sort`.
2. You should implement that function for all of Strings, Ints and Float64s.
3. You must include tests.
4. Update the README to list the new functions.

## Why is the emoji a slice of pizza instead of a pie?

I wanted to pick a name for the project that was short and had an associated emoji. I liked pie, but then I found out that the pie emoji is not fully supported everywhere. I didn't want to change the name of the project to cake, but pizza pie still made sense. I'm not sure if I will change it back to a pie later.
