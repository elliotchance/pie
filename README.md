# 🍕 `github.com/elliotchance/pie`
[![GoDoc](https://godoc.org/github.com/elliotchance/pie?status.svg)](https://godoc.org/github.com/elliotchance/pie)
[![Build Status](https://travis-ci.org/elliotchance/pie.svg?branch=master)](https://travis-ci.org/elliotchance/pie)

**Enjoy a slice!** `pie` is a code generator for dealing with slices that
focuses on type safety, performance and immutability.

- [Quick Start](#quick-start)
- [Functions](#functions)
- [FAQ](#faq)
  * [What are the requirements?](#what-are-the-requirements-)
  * [What are the goals of `pie`?](#what-are-the-goals-of--pie--)
  * [Can I contribute?](#can-i-contribute-)
  * [Why is the emoji a slice of pizza instead of a pie?](#why-is-the-emoji-a-slice-of-pizza-instead-of-a-pie-)

# Quick Start

1. Install `pie`:

```bash
go get -u github.com/elliotchance/pie
```

2. Annotate the types in your source code:

```go
//go:generate pie MyStrings
type Strings []string
```

3. Run `go generate`. This will create a file called `strings_pie.go`. You
should commit this with the rest of your code. Run `go generate` any time you
need to add more types.

4. Usage:

Since these are aliases they can be used interchangeably:

```go
names := Strings{"Bob", "Sally", "John", "Jane"}
shortNames = names.Only(func(s string) bool {
	return len(s) <= 3
})

// shortNames = Strings{"Bob"}
```

Or, more complex operations can be chained:

```go
Strings{"Bob", "Sally", "John", "Jane"}.
	Without(pie.Prefix("J")).
	Transform(strings.ToUpper).
	Last()

// "SALLY"
```

# Functions

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

# FAQ

## What are the requirements?

`pie` supports many Go versions, all the way back to Go 1.8.

## What are the goals of `pie`?

1. **Type safety.** I never want to hit runtime bugs because I could pass in the
wrong type, or perform an invalid type case out the other end.

2. **Performance.** The functions need to be as fast as native Go
implementations otherwise there's no point in this library existing.

3. **Nil-safe.** All of the functions will happily accept nil and treat them as
empty slices. Apart from less possible panics, it makes it easier to chain.

4. **Immutable.** Functions never modify inputs, unlike some built-ins such as
sort.Strings.

## Can I contribute?

Absolutely. Pull requests are always welcome. Your PR must include:

1. Unit tests.
2. Update the README to list the new functions.

## Why is the emoji a slice of pizza instead of a pie?

I wanted to pick a name for the project that was short and had an associated
emoji. I liked pie, but then I found out that the pie emoji is not fully
supported everywhere. I didn't want to change the name of the project to cake,
but pizza pie still made sense. I'm not sure if I will change it back to a pie
later.
