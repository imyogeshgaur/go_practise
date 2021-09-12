package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func setLenghtAndBreadth(r Rectangle) (int,int){
	r.length = 4
	r.breadth = 5
	return r.length, r.breadth

}

func getArea(len int, br int) {
	fmt.Printf("The area of Rectangle is : %d\n", len * br)
}

func getPerimeter(len int, br int) {
	fmt.Printf("The perimeter of Rectangle is : %d", 2*(len + br))
}

func main() {
	rect := Rectangle{}
	l ,b:= setLenghtAndBreadth(rect)
	getArea(l,b)
	getPerimeter(l,b)

}