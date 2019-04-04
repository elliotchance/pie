# 🍕 `github.com/elliotchance/pie` [![GoDoc](https://godoc.org/github.com/elliotchance/pie?status.svg)](https://godoc.org/github.com/elliotchance/pie)

**Enjoy a slice!** `pie` is a utility library for dealing with slices that
focuses on type safety and performance.

It can be used with the Go-style package functions:

```go
names := []string{"Bob", "Sally", "John", "Jane"}
shortNames := pie.StringsOnly(names, func(s string) bool {
	return len(s) <= 3
})

// []string{"Bob"}
```

Or, they can be chained for more complex operations:

```go
pie.Strings{"Bob", "Sally", "John", "Jane"}.
	Without(pie.Prefix("J")).
	Transform(strings.ToUpper).
	Last()

// "SALLY"
```

# Functions

## Slices

| Function    | Description | Strings | Ints  | Float64s |       |
| ----------- | ----------- | :-----: | :---: | :------: | :---: |
| `Average`   | The average (mean) value, or a zeroed value. | | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsAverage) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sAverage)| O(n) |
| `Contains`  | Check if the value exists in the slice. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsContains) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsContains) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sContains)| O(n) |
| `First`     | The first element, or a zeroed value. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsFirst) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsFirst) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sFirst)| O(1) |
| `FirstOr`   | The first element, or a default value. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsFirstOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsFirstOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sFirstOr)| O(1) |
| `Last`      | The last element, or a zeroed value. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsLast) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsLast) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sLast)| O(1) |
| `LastOr`    | The last element, or a default value. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsLastOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsLastOr) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sLastOr)| O(1) |
| `Len`       | Number of elements. | [Yes](https://godoc.org/github.com/elliotchance/pie#Strings.Len) | [Yes](https://godoc.org/github.com/elliotchance/pie#Ints.Len) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64s.Len) | O(1) |
| `Max`       | The maximum value, or a zeroes value. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsMax) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsMax) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sMax)| O(n) |
| `Min`       | The minimum value, or a zeroed value. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsMin) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsMin) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sMin)| O(n) |
| `Only`      | A new slice containing only the elements that returned true from the condition. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsOnly) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsOnly) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sOnly)| O(n) |
| `Sum`       | Sum (total) of all elements. | | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsSum) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sSum)| O(n) |
| `Transform` | A new slice where each element has been transformed. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsTransform) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsTransform) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sTransform)| O(n) |
| `Without`   | A new slice containing only the elements that returned false from the condition. | [Yes](https://godoc.org/github.com/elliotchance/pie#StringsWithout) | [Yes](https://godoc.org/github.com/elliotchance/pie#IntsWithout) | [Yes](https://godoc.org/github.com/elliotchance/pie#Float64sWithout)| O(n) |

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
