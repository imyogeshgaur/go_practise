package main

import "fmt"

func main(){
	var num1 ,num2,num3 int
	fmt.Println("Enter the First Number : ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the Second Number : ")
	fmt.Scanln(&num2)
	fmt.Println("Enter the Second Number : ")
	fmt.Scanln(&num3)
	fmt.Println("The Result of the Numbers is ",num1+num2+num3)
}