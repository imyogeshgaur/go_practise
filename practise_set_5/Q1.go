package main;

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j := 4; j != i; j-- {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
