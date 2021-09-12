package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mutliplication (num int){
	for i := 1; i <= num; i++ {
		fmt.Printf("%d * %d = %d\n", num,i, num*i);
	}
}

func main() {
	numberPointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter the Number : ")

	numberReader, _ := numberPointer.ReadString('\n');

	number, err := strconv.ParseInt(strings.TrimSpace(numberReader),10,64)

	if err == nil {
		newNumber := int(number)
		mutliplication(newNumber)
	}

}