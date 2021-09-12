package main

import "fmt"

func printPattern()  {
	for i := 1; i <= 4; i++  {
		for j := 1; j<=i ; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}


func main()  {
	printPattern()
}