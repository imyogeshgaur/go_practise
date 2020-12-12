package main

import "fmt"

func main(){
	var userNumber,comparedNumber int
	fmt.Printf("Enter the Comparison Number : ")
	fmt.Scanln(&userNumber)
	fmt.Println("Enter the Number To be Compared : ")
	fmt.Scanln(&comparedNumber)
    if(comparedNumber>userNumber){
		fmt.Printf("The Number is Greater !!!")
	}else{
		fmt.Printf("The Number is Smaller!!!")
	}
}