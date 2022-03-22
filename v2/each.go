package pie

// Each is more condensed version of Transform that allows an action to happen
// on each elements and pass the original slice on.
//
//   pie.Each(cars, func (car *Car) {
//       fmt.Printf("Car color is: %s\n", car.Color)
//   })
//
// Pie will not ensure immutability on items passed in so they can be
// manipulated, if you choose to do it this way, for example:
//
//   // Set all car colors to Red.
//   pie.Each(cars, func (car *Car) {
//       car.Color = "Red"
//   })
//
func Each[T any](ss []T, fn func(T)) []T {
	for _, s := range ss {
		fn(s)
	}

	return ss
}
