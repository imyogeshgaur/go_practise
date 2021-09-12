package main

import "fmt"

func main() {
	number := 8;
	sum :=0;

	for i := 1; i <= 10; i++ {
		sum += number * i
	}
	fmt.Printf("The Sum Of Number is %d", sum)
}