package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumOfNaturalNumbers(r int) int {
	sum:=0;
	for i := 1; i <= r; i++ {
		sum += i
	}
	return sum
}

func main() {
	numberPointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter the Range : ");

	numberReader, _ := numberPointer.ReadString('\n');

	number,err := strconv.ParseInt(strings.TrimSpace(numberReader),10,64);

	if err == nil {
		sum := sumOfNaturalNumbers(int(number))

		fmt.Printf("The Sum of Natural number upto %d is %d", number,sum)
	}
}