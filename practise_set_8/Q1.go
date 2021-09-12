package main

import "fmt"

type Employee struct {
	nameOfEmployee  string
	salaryOfEmployee int
}

func setSalry(emp1 Employee) int {
	emp1.salaryOfEmployee = 25000000
	return emp1.salaryOfEmployee;
}
func getSalry(sal int)  {
	fmt.Println("The salary of Employee is : ",sal)
}
func setName(emp1 Employee) string {
	emp1.nameOfEmployee = "Yogesh Gaur"
	return emp1.nameOfEmployee
}
func getName(name string)  {
	fmt.Println("The name of The Employee is : ",name)
}

func main() {
   emp := Employee{}
   name := setName(emp)
   salary := setSalry(emp)
   getName(name)
   getSalry(salary)
   
}