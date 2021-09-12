package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numberPointer := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the Number : ")

	numberReader , _ := numberPointer.ReadString('\n');

	number , err := strconv.ParseInt(strings.TrimSpace(numberReader),10,64);

	newNumber := int(number);

	if err == nil {
		fact := 1;
		for i := newNumber; i>=1; i-- {
			fact *= i;
		}
		fmt.Printf("The Factorial of %d is %d", newNumber,fact);
	}

}