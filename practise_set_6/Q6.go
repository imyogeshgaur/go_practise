package main

import "fmt"

func main(){
	var arr[5] int
	for i := 0; i < 5; i++ {
		fmt.Printf("The Value at %d index is : ",i)
		fmt.Scanln(&arr[i])
	}
	var max = arr[0]
	for i := 0; i < 5; i++ {
		if max < arr[i]{
			max = arr[i]
		}
	}
	fmt.Print("The Max of the Array is : ",max)
}