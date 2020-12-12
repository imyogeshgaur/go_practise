package main

import "fmt"

func main(){
	var sub1,sub2,sub3,CGPA,tot float64
	fmt.Printf("Enter the Marks in Subject1 out of 100 : ")
	fmt.Scanln(&sub1)
	fmt.Printf("Enter the Marks in Subject2 out of 100 : ")
	fmt.Scanln(&sub2)
	fmt.Printf("Enter the Marks in Subject3 out of 100 : ")
	fmt.Scanln(&sub3)
	tot = sub1+sub2+sub3
	CGPA = tot/30
	fmt.Printf("The CGPA odf the student is %.2f ",CGPA)
}