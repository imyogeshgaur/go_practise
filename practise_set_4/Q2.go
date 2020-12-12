package main

import "fmt"

func main(){
	var sub1,sub2,sub3,tot,perc float64
	fmt.Printf("Enter the Marks in Subject1 : ")
	fmt.Scanln(&sub1)
	fmt.Printf("Enter the Marks in Subject2 : ")
	fmt.Scanln(&sub2)
	fmt.Printf("Enter the Marks in Subject3 : ")
	fmt.Scanln(&sub3)
	tot = sub1 + sub2 + sub3
	perc = tot/3
	if tot >= 99 && perc>=40{
		fmt.Printf("The Student is Passed !!")
	}else{
		fmt.Printf("The Student is Fail !!")
	}
}