package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

type carPointers []*car

type car struct {
	Name, Color string
}

func carPointerNameLess(a, b *car) bool {
	return a.Name < b.Name
}

func carPointerColorLess(a, b *car) bool {
	return a.Color < b.Color
}

var carPointersSortCustomTests = []struct {
	ss                  carPointers
	sortedStableByName  carPointers
	sortedStableByColor carPointers
}{
	{
		nil,
		nil,
		nil,
	},
	{
		carPointers{},
		carPointers{},
		carPointers{},
	},
	{
		carPointers{&car{"foo", "red"}},
		carPointers{&car{"foo", "red"}},
		carPointers{&car{"foo", "red"}},
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"foo", "red"}},
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}, &car{"foo", "red"}},
		carPointers{&car{"Baz", "black"}, &car{"foo", "red"}, &car{"bar", "yellow"}},
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"qux", "cyan"}, &car{"foo", "red"}},
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}, &car{"foo", "red"}, &car{"qux", "cyan"}},
		carPointers{&car{"Baz", "black"}, &car{"qux", "cyan"}, &car{"foo", "red"}, &car{"bar", "yellow"}},
	},
	{
		carPointers{&car{"aaa", "yellow"}, &car{"aaa", "black"}, &car{"bbb", "yellow"}, &car{"bbb", "black"}},
		carPointers{&car{"aaa", "yellow"}, &car{"aaa", "black"}, &car{"bbb", "yellow"}, &car{"bbb", "black"}},
		carPointers{&car{"aaa", "black"}, &car{"bbb", "black"}, &car{"aaa", "yellow"}, &car{"bbb", "yellow"}},
	},
}

func TestSortUsing(t *testing.T) {
	isSortedUsing := func(ss carPointers, less func(a, b *car) bool) bool {
		for i := 1; i < len(ss); i++ {
			if less(ss[i], ss[i-1]) {
				return false
			}
		}
		return true
	}

	for _, test := range carPointersSortCustomTests {
		t.Run("", func(t *testing.T) {
			sortedByName := pie.SortUsing(test.ss, carPointerNameLess)
			assert.True(t, isSortedUsing(sortedByName, carPointerNameLess))
			sortedStableByName := pie.SortStableUsing(test.ss, carPointerNameLess)
			assert.Equal(t, test.sortedStableByName, carPointers(sortedStableByName))

			sortedByColor := pie.SortUsing(test.ss, carPointerColorLess)
			assert.True(t, isSortedUsing(sortedByColor, carPointerColorLess))
			sortedStableByColor := pie.SortStableUsing(test.ss, carPointerColorLess)
			assert.Equal(t, test.sortedStableByColor, carPointers(sortedStableByColor))
		})
	}
}
