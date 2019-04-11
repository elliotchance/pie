package pie

//go:generate pie cars
type cars []car

type car struct {
	Name, Color string
}
