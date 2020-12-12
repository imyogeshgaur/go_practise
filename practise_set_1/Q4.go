package main


import "fmt"

func main(){
	var kilo,mile float64
	fmt.Printf("Enter the Distance in Kilometers : ")
	fmt.Scanln(&kilo)
	mile = kilo / 1.609
	fmt.Printf("The Distance in Miles is %.2f",mile)
}