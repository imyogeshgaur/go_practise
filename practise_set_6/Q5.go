package main

import "fmt"

func main(){
	var arr[5] int
	for i := 0; i < 5; i++ {
		fmt.Printf("The Value at %d index is : ",i)
		fmt.Scanln(&arr[i])
	}
	fmt.Print("The Original Array is : ")
	for i := 0; i < 5; i++ {
		fmt.Print(arr[i]," ")
	}
	fmt.Print("Reverse Array is : ")
	for i := 4; i >= 0; i-- {
		fmt.Print(arr[i]," ")
	}
}