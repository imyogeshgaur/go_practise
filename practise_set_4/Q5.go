package main

import "fmt"

func main(){
	var year int
	fmt.Printf("Enter the Year : ")
	fmt.Scanln(&year)
	if year%4 == 0{
		fmt.Println("It is a Leap Year !!!")
	}else{
		fmt.Println("It is not a Leap Year !!")
	}
}