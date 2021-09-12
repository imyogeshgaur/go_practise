package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNaturalNumber(num int) int {
	if num !=0 {
		return num + getNaturalNumber(num-1);
	}
	return 0;
}

func main() {
	numberPointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter the Number : ");

	numberReader, _  := numberPointer.ReadString('\n');

	number, err := strconv.ParseInt(strings.TrimSpace(numberReader),10,64);

	if err == nil {
		newNumber := int(number);
		res := getNaturalNumber(newNumber)
		fmt.Printf("The Total Natural Number in range upto %d is %d", number,res)
	}

}