package pie

import "fmt"

//go:generate pie cars.* carPointers.*
type cars []car
type carPointers []*car

type car struct {
	Name, Color string
}

func (c *car) Equals(c2 *car) bool {
	if c == nil && c2 == nil {
		return true
	}

	if c == nil || c2 == nil {
		return false
	}

	return c.Name == c2.Name
}

func (c *car) String() string {
	return fmt.Sprintf("%s is %s", c.Name, c.Color)
}
