# 🍕 `github.com/elliotchance/pie` [![GoDoc](https://godoc.org/github.com/elliotchance/pie?status.svg)](https://godoc.org/github.com/elliotchance/pie)

**Enjoy a slice!** `pie` is a utility library for dealing with slices that
focuses on type safety, performance and immutability.

- [Quick Start](#quick-start)
- [Functions](#functions)
  * [Slices](#slices)
  * [Conditional](#conditional)
  * [Transforms](#transforms)
- [FAQ](#faq)
  * [How do I use it?](#how-do-i-use-it-)
  * [Why do we need another library for this?](#why-do-we-need-another-library-for-this-)
  * [What are the goals of `pie`?](#what-are-the-goals-of--pie--)
  * [Can I contribute?](#can-i-contribute-)
  * [Why is the emoji a slice of pizza instead of a pie?](#why-is-the-emoji-a-slice-of-pizza-instead-of-a-pie-)
  * [Why does it not have package level functions?](#why-does-it-not-have-package-level-functions-)

# Quick Start

Pie has three basic slice types:

- `Strings` is an alias for `[]string`.
- `Ints` is an alias for `[]int`.
- `Float64s` is an alias for `[]float64`.

Since these are aliases they can be used interchangeably:

```go
names := []string{"Bob", "Sally", "John", "Jane"}
var shortNames []string = pie.Strings(names).Only(func(s string) bool {
	return len(s) <= 3
})

// shortNames = []string{"Bob"}
```

Or, more complex operations can be chained:

```go
pie.Strings{"Bob", "Sally", "John", "Jane"}.
	Without(pie.Prefix("J")).
	Transform(strings.ToUpper).
	Last()

// "SALLY"
```

# Functions

## Slices

| Function     | Description | Strings | Ints  | Float64s | O     |
| ------------ | ----------- | :-----: | :---: | :------: | :---: |
| `AreSorted`  | Check if the slice is already sorted. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.AreSorted) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.AreSorted) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.AreSorted)| n |
| `Average`    | The average (mean) value, or a zeroed value. | | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Average) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Average)| n |
| `Contains`   | Check if the value exists in the slice. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Contains) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Contains) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Contains)| n |
| `First`      | The first element, or a zeroed value. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.First) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.First) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.First)| 1 |
| `FirstOr`    | The first element, or a default value. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.FirstOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.FirstOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.FirstOr)| 1 |
| `JSONString` | The JSON encoded string. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.JSONString) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.JSONString) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.JSONString)| n |
| `Last`       | The last element, or a zeroed value. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Last) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Last) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Last)| 1 |
| `LastOr`     | The last element, or a default value. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.LastOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.LastOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.LastOr)| 1 |
| `Len`        | Number of elements. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Len) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Len) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Len) | 1 |
| `Max`        | The maximum value, or a zeroes value. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Max) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Max) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Max)| n |
| `Min`        | The minimum value, or a zeroed value. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Min) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Min) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Min)| n |
| `Only`       | A new slice containing only the elements that returned true from the condition. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Only) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Only) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Only)| n |
| `Reverse`    | Reverse elements. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Reverse) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Reverse) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Reverse)| n |
| `Sort`       | Return a new sorted slice. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Sort) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Sort) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Sort)| *n⋅log(n)* |
| `Sum`        | Sum (total) of all elements. | | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Sum) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Sum)| n |
| `Transform`  | A new slice where each element has been transformed. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Transform) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Transform) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Transform)| n |
| `Without`    | A new slice containing only the elements that returned false from the condition. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Without) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Without) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Without)| n |

## Conditional

| Function           | Description | Strings | Ints  | Float64s |
| ------------------ | ----------- | :-----: | :---: | :------: |
| `Equal`            | Check if the values are equal. | [Yes](https://godoc.org/github.com/elliotchance/pie#EqualString) | [Yes](https://godoc.org/github.com/elliotchance/pie#EqualInt) | [Yes](https://godoc.org/github.com/elliotchance/pie#EqualFloat64) |
| `GreaterThan`      | Check if the values are greater than. | [Yes](https://godoc.org/github.com/elliotchance/pie#GreaterThanString) | [Yes](https://godoc.org/github.com/elliotchance/pie#EqualInt) | [Yes](https://godoc.org/github.com/elliotchance/pie#GreaterThanFloat64) |
| `GreaterThanEqual` | Check if the values are greater than or equal to. | [Yes](https://godoc.org/github.com/elliotchance/pie#GreaterThanEqualString) | [Yes](https://godoc.org/github.com/elliotchance/pie#GreaterThanEqualInt) | [Yes](https://godoc.org/github.com/elliotchance/pie#GreaterThanEqualFloat64) |
| `LessThan`         | Check if the values are less than. | [Yes](https://godoc.org/github.com/elliotchance/pie#LessThanString) | [Yes](https://godoc.org/github.com/elliotchance/pie#LessThanInt) | [Yes](https://godoc.org/github.com/elliotchance/pie#LessThanFloat64) |
| `LessThanEqual`    | Check if the values are less than or equal to. | [Yes](https://godoc.org/github.com/elliotchance/pie#LessThanEqualString) | [Yes](https://godoc.org/github.com/elliotchance/pie#LessThanEqualInt) | [Yes](https://godoc.org/github.com/elliotchance/pie#LessThanEqualFloat64) |
| `NotEqual`         | Check if the values are not equal. | [Yes](https://godoc.org/github.com/elliotchance/pie#NotEqualString) | [Yes](https://godoc.org/github.com/elliotchance/pie#NotEqualInt) | [Yes](https://godoc.org/github.com/elliotchance/pie#NotEqualFloat64) |
| `Prefix`           | Check if the string starts with another string. | [Yes](https://godoc.org/github.com/elliotchance/pie#Prefix) | | |
| `Suffix`           | Check if the string ends with another string. | [Yes](https://godoc.org/github.com/elliotchance/pie#Suffix) | | |

## Transforms

Some of the functions listed here are part of the Go native libraries, but they
are useful to list.

| Function    | Description | Strings | Ints  | Float64s |
| ----------- | ----------- | :-----: | :---: | :------: |
| `Add`       | Addition.   | | [Yes](https://godoc.org/github.com/elliotchance/pie#AddInt) | [Yes](https://godoc.org/github.com/elliotchance/pie#AddFloat64) |
| `ToLower`   | Convert string to lower case. | [`strings`](https://golang.org/pkg/strings/#ToLower) | | |
| `ToTitle`   | Convert string to title case. | [`strings`](https://golang.org/pkg/strings/#ToTitle) | | |
| `ToUpper`   | Convert string to upper case. | [`strings`](https://golang.org/pkg/strings/#ToUpper) | | |
| `TrimSpace` | Trim leading and trailing whitespace. | [`strings`](https://golang.org/pkg/strings/#TrimSpace) | | |

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

4. **Immutable.** Functions never modify inputs, unlike some built-ins such as
sort.Strings.

There are some downsides with this approach:

1. It won't support all slice types. Sorry, you can use these functions on
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

## Why does it not have package level functions?

It's temping to add package level functions, like `pie.StringsOnly(ss, fn)` as a
shortcut for `pie.Strings(ss).Only(fn)`. In fact older versions did actually
have these. I removed them because:

1. The `pie` types are aliases so there is no special operation to use
`[]string` and `pie.Strings` interchangeably.

2. It created a huge amount of duplicate code that also created lots of extra
tests and more complicated documentation.
