package main

import "fmt"

type Sphere struct {
	radiusOfSphere float64
}

func setter(c Sphere) (float64) {
	c.radiusOfSphere = 5
	return  c.radiusOfSphere
}

func getArea(r float64, p float64)  {
	fmt.Printf("The Area of Sphere is %f\n", 4 * p * r * r)
}

func getVolume(r float64, p float64)  {
	fmt.Printf("The Volume of Sphere is %f", (4/3)* p * r * r * r)
}

func main() {
	cy := Sphere{}
	radius := setter(cy)
	pi := 3.14
	getArea(radius,pi)
	getVolume(radius,pi)
}
