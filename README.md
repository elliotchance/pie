# 🍕 `github.com/elliotchance/pie`
[![GoDoc](https://godoc.org/github.com/elliotchance/pie?status.svg)](https://godoc.org/github.com/elliotchance/pie)
[![Build Status](https://travis-ci.org/elliotchance/pie.svg?branch=master)](https://travis-ci.org/elliotchance/pie)
[![codecov](https://codecov.io/gh/elliotchance/pie/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotchance/pie)

**Enjoy a slice!** `pie` is a library of utility functions for common operations
on slices and maps.

- [Quick Start](#quick-start)
- [FAQ](#faq)
  * [What are the requirements?](#what-are-the-requirements-)
  * [What are the goals of `pie`?](#what-are-the-goals-of--pie--)
  * [How do I contribute a function?](#how-do-i-contribute-a-function-)
  * [Why is the emoji a slice of pizza instead of a pie?](#why-is-the-emoji-a-slice-of-pizza-instead-of-a-pie-)

# Quick Start

If you are using (or require) Go 1.17 or below, you will have to
[use v1](https://github.com/elliotchance/pie/v1).

`pie` can be used in two ways, the first is to use the regular
[parameterized functions](https://go.googlesource.com/proposal/+/master/design/15292/2013-12-type-params.md):

[Run this program](https://go.dev/play/p/qYaBXPRs3Nk)

```go
package main

import (
    "fmt"
    "strings"

    "github.com/elliotchance/pie/v2"
)

func main() {
    names := pie.FilterNot([]string{"Bob", "Sally", "John", "Jane"},
        func(name string) bool {
            return strings.HasPrefix(name, "J")
        })

    fmt.Println(names) // "[Bob Sally]"
}
```

Or, if you need to chain multiple operations you can use one of:

- [`pie.Of`](https://pkg.go.dev/github.com/elliotchance/pie/v2#Of) - works with any element type, but functions are limited.
- [`pie.OfOrdered`](https://pkg.go.dev/github.com/elliotchance/pie/v2#OfOrdered) - only works with numbers and strings, but has more functions.
- [`pie.OfNumeric`](https://pkg.go.dev/github.com/elliotchance/pie/v2#OfNumeric) - only works with numbers, but has all functions.

[Run this program](https://go.dev/play/p/cDLBYzKJ9ld)

```go
package main

import (
    "fmt"
    "strings"

    "github.com/elliotchance/pie/v2"
)

func main() {
    name := pie.Of([]string{"Bob", "Sally", "John", "Jane"}).
        FilterNot(func(name string) bool {
            return strings.HasPrefix(name, "J")
        }).
        Map(strings.ToUpper).
        LastOr("")

    fmt.Println(name) // "SALLY"
}
```

You can find the
[full documentation here](https://pkg.go.dev/github.com/elliotchance/pie/v2).

# FAQ

## What are the requirements?

`pie` v2 only supports Go 1.18+. If you have an older version you can
[use v1](https://github.com/elliotchance/pie/v1).

## What are the goals of `pie`?

1. **Type safety.** I never want to hit runtime bugs because I could pass in the
wrong type, or perform an invalid type case out the other end.

2. **Performance.** The functions need to be as fast as native Go
implementations otherwise there's no point in this library existing.

3. **Nil-safe.** All of the functions will happily accept nil and treat them as
empty slices. Apart from less possible panics, it makes it easier to chain.

4. **Immutable.** Functions never modify inputs (except in cases where it would
be illogical), unlike some built-ins such as `sort.Strings`.

## How do I contribute a function?

Pull requests are always welcome.

Here is a comprehensive list of steps to follow to add a new function:

1. Create a new file for your function (tip: copy an existing file can be
quicker). Add your implmentation and comment.

2. Create appropriate tests.

3. If your function accepts a slice, it should also be added to the `OfSlice`
API (see `of.go`).

## Why is the emoji a slice of pizza instead of a pie?

I wanted to pick a name for the project that was short and had an associated
emoji. I liked pie, but then I found out that the pie emoji is not fully
supported everywhere. I didn't want to change the name of the project to cake,
but pizza pie still made sense. I'm not sure if I will change it back to a pie
later.
