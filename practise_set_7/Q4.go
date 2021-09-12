package main

import "fmt"

func printPattern() {
	for i := 0; i < 4; i++ {
		for j := i; j < 4; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}

func main() {
	printPattern()
}