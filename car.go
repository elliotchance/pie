package main

//go:generate pie Cars
type Cars []Car

type Car struct {
	Name, Color string
}
