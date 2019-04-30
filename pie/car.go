package pie

//go:generate pie cars.* carPointers.*
type cars []car
type carPointers []*car

type car struct {
	Name, Color string
}
