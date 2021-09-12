package main

import "fmt"

type Square struct{
	sideOfSquare int
}

func setSide(s Square) int {
	s.sideOfSquare = 4;
	return s.sideOfSquare
}

func getArea(s int) {
	fmt.Printf("The Area Of Square is %d\n", s * s)
}

func getPerimeter(s int)  {
	fmt.Printf("The Perimeter of Square is %d\n", 4 * s)
}


func main(){
	sq := Square{}
	side := setSide(sq)
	getArea(side)
	getPerimeter(side)
}