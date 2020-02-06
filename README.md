# 🍕 `github.com/elliotchance/pie`
[![GoDoc](https://godoc.org/github.com/elliotchance/pie?status.svg)](https://godoc.org/github.com/elliotchance/pie)
[![Build Status](https://travis-ci.org/elliotchance/pie.svg?branch=master)](https://travis-ci.org/elliotchance/pie)
[![codecov](https://codecov.io/gh/elliotchance/pie/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotchance/pie)

**Enjoy a slice!** `pie` is a code generator for dealing with slices that
focuses on type safety, performance and immutability.

- [Quick Start](#quick-start)
  * [Install/Update](#install-update)
  * [Built-in Types](#built-in-types)
  * [Custom Types](#custom-types)
  * [Custom Equality](#custom-equality)
  * [Custom Stringer](#custom-stringer)
  * [Limiting Functions Generated](#limiting-functions-generated)
- [Functions](#functions)
- [FAQ](#faq)
  * [What are the requirements?](#what-are-the-requirements-)
  * [What are the goals of `pie`?](#what-are-the-goals-of--pie--)
  * [How do I contribute a function?](#how-do-i-contribute-a-function-)
  * [Why is the emoji a slice of pizza instead of a pie?](#why-is-the-emoji-a-slice-of-pizza-instead-of-a-pie-)
  * [How do I exclude generated files from code coverage?](#how-do-i-exclude-generated-files-from-code-coverage-)

# Quick Start

## Install/Update

```bash
go get -u github.com/elliotchance/pie
```

## Built-in Types

`pie` ships with some slice types ready to go (pun intended). These include:

- `type`[`Strings`](https://godoc.org/github.com/elliotchance/pie/pie#Strings)`[]string`
- `type`[`Float64s`](https://godoc.org/github.com/elliotchance/pie/pie#Float64s)`[]float64`
- `type`[`Ints`](https://godoc.org/github.com/elliotchance/pie/pie#Ints)`[]int`

These can be used without needing `go generate`. For example:

```go
package main

import (
    "fmt"
    "strings"

    "github.com/elliotchance/pie/pie"
)

func main() {
    name := pie.Strings{"Bob", "Sally", "John", "Jane"}.
        FilterNot(func (name string) bool {
            return strings.HasPrefix(name, "J")
        }).
        Map(strings.ToUpper).
        Last()

    fmt.Println(name) // "SALLY"
}
```

## Custom Types

Annotate the slice type in your source code:

```go
type Car struct {
    Name, Color string
}

//go:generate pie Cars.*
type Cars []Car
```

Run `go generate`. This will create a file called `cars_pie.go`. You should
commit this with the rest of your code. Run `go generate` any time you need to
add more types.

Now you can use the slices:

```go
cars := Cars{
    {"Bob", "blue"},
    {"Sally", "green"},
    {"John", "red"},
    {"Jane", "red"},
}

redCars := cars.Filter(func(car Car) bool {
    return car.Color == "red"
})

// redCars = Cars{{"John", "red"}, {"Jane", "red"}}
```

Or, more complex operations can be chained:

```go
cars.FilterNot(func (car Car) {
        return strings.HasPrefix(car.Name, "J")
    }).
    Map(func (car Car) Car {
        car.Name = strings.ToUpper(car.Name)

        return car
    }).
    Last()

// Car{"SALLY", "green"}
```

## Custom Equality

Some functions that compare elements, such as Contains will use the following
method if it is available on the element type:

```go
func (a ElementType) Equals(b ElementType) bool
```

The `ElementType` must be the same for the receiver and argument and it must
return a bool. Be careful to create the function on the pointer or non-pointer
type that is used by the slice.

Here is a minimal example:

```go
type Car struct {
	Name, Color string
}

type Cars []*Car // ElementType is *Car

func (c *Car) Equals(c2 *Car) bool {
	return c.Name == c2.Name
}
```

## Custom Stringer

Some functions that need elements to be represented as strings, such as
`Strings()` will try to use the `fmt.Stringer` interface. If it's not available
then it will fallback to:

```go
fmt.Sprintf("%v", element)
```

## Limiting Functions Generated

The `.*` can be used to generate all functions. This is easy to get going but
creates a lot of unused code. You can limit the functions generated by chaining
the function names with a dot syntax, like:

```go
//go:generate pie myInts.Average.Sum myStrings.Filter
```

This will only generate `myInts.Average`, `myInts.Sum` and `myStrings.Filter`.

# Functions

Below is a summary of the available functions.

The letters in brackets indicate:

- **E**: The function will use the `Equals` method if it is available. See
*Custom Equality*.

- **S**: The function will use the `String` method if it is available. See
*Custom Stringer*.

| Function                                | String | Number | Struct | Maps | Big-O    | Description |
| --------------------------------------- | :----: | :----: | :----: | :--: | :------: | ----------- |
| [`Abs`](#abs)                           |        | ✓      |        |      | n        | Abs is a function which returns the absolute value of all the elements in the slice.  |
| [`All`](#all)                           | ✓      | ✓      | ✓      |      | n        | All will return true if all callbacks return true. It follows the same logic as the all() function in Python.  |
| [`Any`](#any)                           | ✓      | ✓      | ✓      |      | n        | Any will return true if any callbacks return true. It follows the same logic as the any() function in Python.  |
| [`Append`](#append)                     | ✓      | ✓      | ✓      |      | n        | Append will return a new slice with the elements appended to the end.  |
| [`AreSorted`](#aresorted)               | ✓      | ✓      |        |      | n        | AreSorted will return true if the slice is already sorted. It is a wrapper for sort.SliceTypeAreSorted.  |
| [`AreUnique`](#areunique)               | ✓      | ✓      |        |      | n        | AreUnique will return true if the slice contains elements that are all different (unique) from each other.  |
| [`Average`](#average)                   |        | ✓      |        |      | n        | Average is the average of all of the elements, or zero if there are no elements.  |
| [`Bottom`](#bottom)                     | ✓      | ✓      | ✓      |      | n        | Bottom will return n elements from bottom  |
| [`Contains`](#contains) (E)             | ✓      | ✓      | ✓      |      | n        | Contains returns true if the element exists in the slice.  |
| [`Diff`](#diff) (E)                     | ✓      | ✓      | ✓      |      | n²       | Diff returns the elements that needs to be added or removed from the first slice to have the same elements in the second slice.  |
| [`DropTop`](#droptop)                   | ✓      | ✓      | ✓      |      | n        | DropTop will return the rest slice after dropping the top n elements if the slice has less elements then n that'll return empty slice if n < 0 it'll return empty slice.  |
| [`Each`](#each)                         | ✓      | ✓      | ✓      |      | n        | Each is more condensed version of Transform that allows an action to happen on each elements and pass the original slice on.  |
| [`Equals`](#equals) (E)                 | ✓      | ✓      | ✓      |      | n        | Equals compare elements from the start to the end,  |
| [`Extend`](#extend)                     | ✓      | ✓      | ✓      |      | n        | Extend will return a new slice with the slices of elements appended to the end.  |
| [`Filter`](#filter)                     | ✓      | ✓      | ✓      |      | n        | Filter will return a new slice containing only the elements that return true from the condition. The returned slice may contain zero elements (nil).  |
| [`FilterNot`](#filternot)               | ✓      | ✓      | ✓      |      | n        | FilterNot works the same as Filter, with a negated condition. That is, it will return a new slice only containing the elements that returned false from the condition. The returned slice may contain zero elements (nil).  |
| [`FindFirstUsing`](#findfirstusing)     | ✓      | ✓      | ✓      |      | n        | FindFirstUsing will return the index of the first element when the callback returns true or -1 if no element is found. It follows the same logic as the findIndex() function in Javascript.  |
| [`First`](#first)                       | ✓      | ✓      | ✓      |      | 1        | First returns the first element, or zero. Also see FirstOr().  |
| [`FirstOr`](#firstor)                   | ✓      | ✓      | ✓      |      | 1        | FirstOr returns the first element or a default value if there are no elements.  |
| [`Float64s`](#float64s) (S)             | ✓      | ✓      | ✓      |      | n        | Float64s transforms each element to a float64.  |
| [`Group`](#group)                       | ✓      | ✓      |        |      | n        | Group returns a map of the value with an individual count.  |
| [`Intersect`](#intersect)               | ✓      | ✓      |        |      | n        | Intersect returns items that exist in all lists.  |
| [`Insert`](#insert)                     | ✓      | ✓      | ✓      |      | n        | Insert a value at an index  |
| [`Ints`](#ints) (S)                     | ✓      | ✓      | ✓      |      | n        | Ints transforms each element to an integer.  |
| [`Join`](#join) (S)                     | ✓      | ✓      | ✓      |      | n        | Join returns a string from joining each of the elements.  |
| [`JSONBytes`](#jsonbytes)               | ✓      | ✓      | ✓      |      | n        | JSONBytes returns the JSON encoded array as bytes.  |
| [`JSONBytesIndent`](#jsonbytesindent)   | ✓      | ✓      | ✓      |      | n        | JSONBytesIndent returns the JSON encoded array as bytes with indent applied.  |
| [`JSONString`](#jsonstring)             | ✓      | ✓      | ✓      |      | n        | JSONString returns the JSON encoded array as a string.  |
| [`JSONStringIndent`](#jsonstringindent) | ✓      | ✓      | ✓      |      | n        | JSONStringIndent returns the JSON encoded array as a string with indent applied.  |
| [`Keys`](#keys)                         |        |        |        | ✓    | n        | Keys returns the keys in the map. All of the items will be unique.  |
| [`Last`](#last)                         | ✓      | ✓      | ✓      |      | 1        | Last returns the last element, or zero. Also see LastOr().  |
| [`LastOr`](#lastor)                     | ✓      | ✓      | ✓      |      | 1        | LastOr returns the last element or a default value if there are no elements.  |
| [`Len`](#len)                           | ✓      | ✓      | ✓      |      | 1        | Len returns the number of elements.  |
| [`Map`](#map)                           | ✓      | ✓      | ✓      |      | n        | Map will return a new slice where each element has been mapped (transformed). The number of elements returned will always be the same as the input.  |
| [`Max`](#max)                           | ✓      | ✓      |        |      | n        | Max is the maximum value, or zero.  |
| [`Median`](#median)                     |        | ✓      |        |      | n        | Median returns the value separating the higher half from the lower half of a data sample.  |
| [`Min`](#min)                           | ✓      | ✓      |        |      | n        | Min is the minimum value, or zero.  |
| [`Mode`](#mode)                         | ✓      | ✓      | ✓      |      | n        | Mode returns a new slice containing the most frequently occuring values.  |
| [`Pop`](#pop)                           | ✓      | ✓      | ✓      |      | n        | Pop the first element of the slice  |
| [`Product`](#product)                   |        | ✓      |        |      | n        | Product is the product of all of the elements.  |
| [`Random`](#random)                     | ✓      | ✓      | ✓      |      | 1        | Random returns a random element by your rand.Source, or zero  |
| [`Reduce`](#reduce)                     | ✓      | ✓      |        |      | n        | Reduce continually applies the provided function over the slice. Reducing the elements to a single value.  |
| [`Reverse`](#reverse)                   | ✓      | ✓      | ✓      |      | n        | Reverse returns a new copy of the slice with the elements ordered in reverse. This is useful when combined with Sort to get a descending sort order:  |
| [`Send`](#send)                         | ✓      | ✓      | ✓      |      | n        | Send sends elements to channel in normal act it sends all elements but if func canceled it can be less  |
| [`Sequence`](#sequence)                 |        | ✓      |        |      | n        | Sequence generates all numbers in range or returns nil if params invalid  |
| [`SequenceUsing`](#sequenceusing)       | ✓      | ✓      | ✓      |      | n        | SequenceUsing generates slice in range using creator function  |
| [`Shift`](#shift)                       | ✓      | ✓      | ✓      |      | n        | Shift will return two values: the shifted value and the rest slice.  |
| [`Shuffle`](#shuffle)                   | ✓      | ✓      | ✓      |      | n        | Shuffle returns shuffled slice by your rand.Source  |
| [`Sort`](#sort)                         | ✓      | ✓      |        |      | n⋅log(n) | Sort works similar to sort.SliceType(). However, unlike sort.SliceType the slice returned will be reallocated as to not modify the input slice.  |
| [`SortStableUsing`](#sortstableusing)   | ✓      |        | ✓      |      | n⋅log(n) | SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the slice returned will be reallocated as to not modify the input slice.  |
| [`SortUsing`](#sortusing)               | ✓      |        | ✓      |      | n⋅log(n) | SortUsing works similar to sort.Slice. However, unlike sort.Slice the slice returned will be reallocated as to not modify the input slice.  |
| [`Strings`](#strings) (S)               | ✓      | ✓      | ✓      |      | n        | Strings transforms each element to a string.  |
| [`SubSlice`](#subslice)                 | ✓      | ✓      | ✓      |      | n        | SubSlice will return the subSlice from start to end(excluded)  |
| [`Sum`](#sum)                           |        | ✓      |        |      | n        | Sum is the sum of all of the elements.  |
| [`Top`](#top)                           | ✓      | ✓      | ✓      |      | n        | Top will return n elements from head of the slice if the slice has less elements then n that'll return all elements if n < 0 it'll return empty slice.  |
| [`StringsUsing`](#stringsusing)         | ✓      | ✓      | ✓      |      | n        | StringsUsing transforms each element to a string.  |
| [`Unique`](#unique)                     | ✓      | ✓      |        |      | n        | Unique returns a new slice with all of the unique values.  |
| [`Unshift`](#unshift)                   | ✓      | ✓      | ✓      |      | n        | Unshift adds one or more elements to the beginning of the slice and returns the new slice.  |
| [`Values`](#values)                     |        |        |        | ✓    | n        | Values returns the values in the map.  |

## Abs

Abs is a function which returns the absolute value of all the
elements in the slice.


## All

All will return true if all callbacks return true. It follows the same logic
as the all() function in Python.


If the list is empty then true is always returned.


## Any

Any will return true if any callbacks return true. It follows the same logic
as the any() function in Python.


If the list is empty then false is always returned.


## Append

Append will return a new slice with the elements appended to the end.


It is acceptable to provide zero arguments.


## AreSorted

AreSorted will return true if the slice is already sorted. It is a wrapper
for sort.SliceTypeAreSorted.


## AreUnique

AreUnique will return true if the slice contains elements that are all
different (unique) from each other.


## Average

Average is the average of all of the elements, or zero if there are no
elements.


## Bottom

Bottom will return n elements from bottom


that means that elements is taken from the end of the slice
for this [1,2,3] slice with n == 2 will be returned [3,2]
if the slice has less elements then n that'll return all elements
if n < 0 it'll return empty slice.


## Contains

Contains returns true if the element exists in the slice.


When using slices of pointers it will only compare by address, not value.


## Diff

Diff returns the elements that needs to be added or removed from the first
slice to have the same elements in the second slice.


The order of elements is not taken into consideration, so the slices are
treated sets that allow duplicate items.


The added and removed returned may be blank respectively, or contain upto as
many elements that exists in the largest slice.


## DropTop

DropTop will return the rest slice after dropping the top n elements
if the slice has less elements then n that'll return empty slice
if n < 0 it'll return empty slice.


## Each

Each is more condensed version of Transform that allows an action to happen
on each elements and pass the original slice on.


```go
cars.Each(func (car *Car) {
    fmt.Printf("Car color is: %s\n", car.Color)
})

```

Pie will not ensure immutability on items passed in so they can be
manipulated, if you choose to do it this way, for example:


```go
// Set all car colors to Red.
cars.Each(func (car *Car) {
    car.Color = "Red"
})

```



## Equals

Equals compare elements from the start to the end,


if they are the same is considered the slices are equal if all elements are the same is considered the slices are equal
if each slice == nil is considered that they're equal


if element realizes Equals interface it uses that method, in other way uses default compare


## Extend

Extend will return a new slice with the slices of elements appended to the
end.


It is acceptable to provide zero arguments.


## Filter

Filter will return a new slice containing only the elements that return
true from the condition. The returned slice may contain zero elements (nil).


FilterNot works in the opposite way of Filter.


## FilterNot

FilterNot works the same as Filter, with a negated condition. That is, it will
return a new slice only containing the elements that returned false from the
condition. The returned slice may contain zero elements (nil).


## FindFirstUsing

FindFirstUsing will return the index of the first element when the callback returns true or -1 if no element is found.
It follows the same logic as the findIndex() function in Javascript.


If the list is empty then -1 is always returned.


## First

First returns the first element, or zero. Also see FirstOr().


## FirstOr

FirstOr returns the first element or a default value if there are no
elements.


## Float64s

Float64s transforms each element to a float64.


## Group

Group returns a map of the value with an individual count.




## Intersect

Intersect returns items that exist in all lists.


It returns slice without any duplicates.
If zero slice arguments are provided, then nil is returned.


## Insert

Insert a value at an index


## Ints

Ints transforms each element to an integer.


## Join

Join returns a string from joining each of the elements.


## JSONBytes

JSONBytes returns the JSON encoded array as bytes.


One important thing to note is that it will treat a nil slice as an empty
slice to ensure that the JSON value return is always an array.


## JSONBytesIndent

JSONBytesIndent returns the JSON encoded array as bytes with indent applied.


One important thing to note is that it will treat a nil slice as an empty
slice to ensure that the JSON value return is always an array. See
json.MarshalIndent for details.


## JSONString

JSONString returns the JSON encoded array as a string.


One important thing to note is that it will treat a nil slice as an empty
slice to ensure that the JSON value return is always an array.


## JSONStringIndent

JSONStringIndent returns the JSON encoded array as a string with indent applied.


One important thing to note is that it will treat a nil slice as an empty
slice to ensure that the JSON value return is always an array. See
json.MarshalIndent for details.


## Keys

Keys returns the keys in the map. All of the items will be unique.


Due to Go's randomization of iterating maps the order is not deterministic.


## Last

Last returns the last element, or zero. Also see LastOr().


## LastOr

LastOr returns the last element or a default value if there are no elements.


## Len

Len returns the number of elements.


## Map

Map will return a new slice where each element has been mapped (transformed).
The number of elements returned will always be the same as the input.


Be careful when using this with slices of pointers. If you modify the input
value it will affect the original slice. Be sure to return a new allocated
object or deep copy the existing one.


## Max

Max is the maximum value, or zero.


## Median

Median returns the value separating the higher half from the lower half of a
data sample.


Zero is returned if there are no elements in the slice.


If the number of elements is even, then the ElementType mean of the two "median values"
is returned.


## Min

Min is the minimum value, or zero.


## Mode

Mode returns a new slice containing the most frequently occuring values.


The number of items returned may be the same as the input or less. It will
never return zero items unless the input slice has zero items.


## Pop

Pop the first element of the slice


Usage Example:


```go
type knownGreetings []string
greetings := knownGreetings{"ciao", "hello", "hola"}
for greeting := greetings.Pop(); greeting != nil; greeting = greetings.Pop() {
    fmt.Println(*greeting)
}

```

## Product

Product is the product of all of the elements.


## Random

Random returns a random element by your rand.Source, or zero


## Reduce

Reduce continually applies the provided function
over the slice. Reducing the elements to a single value.


Returns a zero value of ElementType if there are no elements in the slice. It will panic if the reducer is nil and the slice has more than one element (required to invoke reduce).
Otherwise returns result of applying reducer from left to right.


## Reverse

Reverse returns a new copy of the slice with the elements ordered in reverse.
This is useful when combined with Sort to get a descending sort order:


```go
ss.Sort().Reverse()

```



## Send

Send sends elements to channel
in normal act it sends all elements but if func canceled it can be less


it locks execution of gorutine
it doesn't close channel after work
returns sended elements if len(this) != len(old) considered func was canceled


## Sequence

Sequence generates all numbers in range or returns nil if params invalid


There are 3 variations to generate:
		1. [0, n).
	2. [min, max).
	3. [min, max) with step.


if len(params) == 1 considered that will be returned slice between 0 and n,
where n is the first param, [0, n).
if len(params) == 2 considered that will be returned slice between min and max,
where min is the first param, max is the second, [min, max).
if len(params) > 2 considered that will be returned slice between min and max with step,
where min is the first param, max is the second, step is the third one, [min, max) with step,
others params will be ignored


## SequenceUsing

SequenceUsing generates slice in range using creator function


There are 3 variations to generate:
		1. [0, n).
	2. [min, max).
	3. [min, max) with step.


if len(params) == 1 considered that will be returned slice between 0 and n,
where n is the first param, [0, n).
if len(params) == 2 considered that will be returned slice between min and max,
where min is the first param, max is the second, [min, max).
if len(params) > 2 considered that will be returned slice between min and max with step,
where min is the first param, max is the second, step is the third one, [min, max) with step,
others params will be ignored


## Shift

Shift will return two values: the shifted value and the rest slice.


## Shuffle

Shuffle returns shuffled slice by your rand.Source


## Sort

Sort works similar to sort.SliceType(). However, unlike sort.SliceType the
slice returned will be reallocated as to not modify the input slice.


See Reverse() and AreSorted().


## SortStableUsing

SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
slice returned will be reallocated as to not modify the input slice.


## SortUsing

SortUsing works similar to sort.Slice. However, unlike sort.Slice the
slice returned will be reallocated as to not modify the input slice.


## Strings

Strings transforms each element to a string.


If the element type implements fmt.Stringer it will be used. Otherwise it
will fallback to the result of:


```go
fmt.Sprintf("%v")

```



## SubSlice

SubSlice will return the subSlice from start to end(excluded)


Condition 1: If start < 0 or end < 0, nil is returned.
Condition 2: If start >= end, nil is returned.
Condition 3: Return all elements that exist in the range provided,
if start or end is out of bounds, zero items will be placed.


## Sum

Sum is the sum of all of the elements.


## Top

Top will return n elements from head of the slice
if the slice has less elements then n that'll return all elements
if n < 0 it'll return empty slice.


## StringsUsing

StringsUsing transforms each element to a string.


## Unique

Unique returns a new slice with all of the unique values.


The items will be returned in a randomized order, even with the same input.


The number of items returned may be the same as the input or less. It will
never return zero items unless then input slice has zero items.


A slice with zero elements is considered to be unique.


See AreUnique().


## Unshift

Unshift adds one or more elements to the beginning of the slice
and returns the new slice.


## Values

Values returns the values in the map.


Due to Go's randomization of iterating maps the order is not deterministic.


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
`sort.Strings`.

## How do I contribute a function?

Pull requests are always welcome.

Here is a comprehensive list of steps to follow to add a new function:

1. Create a new file in the `functions/` directory. The file should be named the
same as the function. You must include documentation for your function.

2. Update `functions/main.go` to register the new function by adding an entry to
`Functions`. Make sure you choose the correct `For` value that is appropriate
for your function.

3. Run `go generate ./... && go install && go generate ./...`. The first
`generate` is to create the pie templates, `install` will update your binary for
the annotations and the second `generate` will use the newly created templates
to update the generated code for the internal types. If you encounter errors
with your code you can safely rerun the command above.

4. If you chose `ForAll` or `ForStructs`, then you must add unit tests to
`pie/carpointers_test.go` and `pie/cars_test.go`.

5. If you chose `ForAll`, `ForNumbersAndStrings` or `ForNumbers`, then you must
add unit tests to `pie/float64s_test.go` and `pie/ints_test.go`.

6. If you chose `ForAll` or `ForStrings`, then you must add unit tests to
`pie/strings_test.go`.

7. If you chose `ForMaps`, then you must add unit tests to `pie/currencies.go`.

## Why is the emoji a slice of pizza instead of a pie?

I wanted to pick a name for the project that was short and had an associated
emoji. I liked pie, but then I found out that the pie emoji is not fully
supported everywhere. I didn't want to change the name of the project to cake,
but pizza pie still made sense. I'm not sure if I will change it back to a pie
later.

## How do I exclude generated files from code coverage?

Go does not provide a way to exclude code coverage from specific files or lines
with comments. However, you can remove the `_pie.go` files from the code
coverage report before it is published:

```bash
go test -race -coverprofile=coverage.txt -covermode=atomic
sed -i '/_pie\.go/d' ./coverage.txt
```

If you are running on macOS, you will need a slightly different syntax for
`sed`:

```bash
sed -i '' /_pie\.go/d' ./coverage.txt
```
