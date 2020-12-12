package main

import "fmt"

func main(){
	var arr[5] float64
	var sum float64
	sum =0
    for i := 0; i < 5; i++ {
		fmt.Print("Enter the Number : ")
		fmt.Scanln(&arr[i])
		sum += arr[i]
	}
	fmt.Printf("The Array Created is : ")
	for i := 0; i < 5; i++ {
		fmt.Print(arr[i]," ")
	}
	fmt.Print("The sum of the Numbers is ",sum)
}