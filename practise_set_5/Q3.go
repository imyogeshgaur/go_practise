package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numberPointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter the Number : ")

	numberReader,_ := numberPointer.ReadString('\n');

	number ,err := strconv.ParseInt(strings.TrimSpace(numberReader),10,64);

	newNumber := int(number)

	fmt.Printf("The Table of %d is : \n", newNumber)

	if err == nil {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d * %d = %d\n",number,i,newNumber*i);
		}
	}

}