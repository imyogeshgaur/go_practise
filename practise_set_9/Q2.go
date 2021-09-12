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

func getArea(r float64,h float64, p float64)  {
	fmt.Printf("The Area of Cyllinder is %f\n", 2* p * r * (r +h))
}

func getVolume(r float64,h float64, p float64)  {
	fmt.Printf("The Volume of Cyllinder is %f", p * r * r * h)
}

func main() {
	cy := Cyllinder{}
	radius, height := setter(cy)
	pi := 3.14
	getArea(radius,height,pi)
	getVolume(radius,height,pi)
}
