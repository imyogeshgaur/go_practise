package main

import "fmt"

func main()  {
	number :=10;
	for i := 10; i >=1; i-- {
		fmt.Printf("%d * %d = %d\n", number,i,number * i)
	}

}