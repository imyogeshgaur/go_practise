package main

import (
	"fmt"
)

type Circle struct{
	radiusOfCircle float64
}

func setSide(s Circle) float64 {
	s.radiusOfCircle = 4;
	return s.radiusOfCircle
}

func getArea(r float64, p float64) {
	fmt.Printf("The Area Of Circle is %f\n", p * r * r )
}

func getPerimeter(r float64, p float64)  {
	fmt.Printf("The Perimeter of Circle is %f\n", 2 * p *r)
}


func main(){
	cr := Circle{}
	pi := 3.14
	radius := setSide(cr)
	getArea(radius,pi)
	getPerimeter(radius,pi)
}