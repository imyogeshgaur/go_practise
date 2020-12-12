package main

import "fmt"

func main(){
	var arr[5] int
	var numberToSearch int
	for i := 0; i < 5; i++ {
		fmt.Print("Enter the integer at index ",i," : ")
		fmt.Scanln(&arr[i])
	}
	fmt.Print("Enter the Number To be Searched for : ")
	fmt.Scanln(&numberToSearch)
	var flag = 0
	for i := 0; i < 5; i++ {
		if arr[i] == numberToSearch{
			flag = 1
			break
		}
	}
	if(flag==1){
		fmt.Print("The Number is Present in the Array !!")
		}else{
		fmt.Print("The Number is not Present in the Array !!")
	}
}