package main

import "fmt"

func main(){
	var name string
	fmt.Printf("Enter Your name : ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s have a Good Day !!!!",name)
}