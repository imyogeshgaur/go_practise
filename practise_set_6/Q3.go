package main

import "fmt"

func main(){
	var marks[5] int
	var averageMarks int
	var sum int
	sum = 0 
	for i := 0; i < 5; i++ {
		fmt.Print("The marks of Student ",i," is : ")
		fmt.Scanln(&marks[i])
		sum += marks[i]
	}
	averageMarks = sum /5
	fmt.Print("The Average Marks of the Student is : ",averageMarks)
}