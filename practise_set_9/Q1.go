package main

import "fmt"

type Cyllinder struct {
	heightOfCyllinder float64
	radiusOfCyllinder float64
}

func setter(c Cyllinder) (float64, float64) {
	c.heightOfCyllinder = 4
	c.radiusOfCyllinder = 5
	return c.heightOfCyllinder, c.radiusOfCyllinder
}

func getter(r float64, h float64) {
	fmt.Printf("The Radius Of Cyllinder is %f and height of Cyllinder is %f ", r,h)
}

func main() {
	cy := Cyllinder{}
	radius, height := setter(cy)
	getter(radius,height)
}
