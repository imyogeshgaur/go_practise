package main;

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func getSumOfFibo(num int) int{
	if num == 0 {
		return 0;
    }else if num == 1{
		return 1;
	}else{
		return (getSumOfFibo(num-1) + getSumOfFibo(num-2));
	}

}


func main()  {
	termPointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter the Number : ");

	termReader, _  := termPointer.ReadString('\n');

	term, err := strconv.ParseInt(strings.TrimSpace(termReader),10,64);

	if err == nil {
		newTerm := int(term);
		res := getSumOfFibo(newTerm)
		fmt.Printf("The Sum upto %d is %d", term,res)
    }
}