package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	numberPointer := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the Number : ");
	numberReder , _ := numberPointer.ReadString('\n');

	number,err  := strconv.ParseInt(strings.TrimSpace(numberReder),10,64);

	var sum int = 0;
	if err == nil {
		for i := 0; i <= int(number); i++ {
			if(i % 2  == 0){
				sum += i;
			}
		}
		fmt.Printf("The Sum of the Even Numbers is %d ", sum)
	}	

}