package main

import "fmt"

func main(){
	var arr[5] int
	for i := 0; i < 5; i++ {
		fmt.Printf("The Value at %d index is : ",i)
		fmt.Scanln(&arr[i])
	}
	var min = arr[0]
	for i := 0; i < 5; i++ {
		if min > arr[i]{
			min = arr[i]
		}
	}
	fmt.Print("The Min of the Array is : ",min)
}