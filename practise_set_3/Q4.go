package main

import (
	"fmt"
	"strings"
)

func main(){
	var originalString string
	fmt.Printf("Enter the String : ")
	fmt.Scanln(&originalString)
	if strings.ContainsAny(originalString,"  ") {
		fmt.Printf("Double Space Detected !!!")
		} else if strings.ContainsAny(originalString,"   ") {
			fmt.Printf("Triple Space Detected !!!")
	}else {
		fmt.Printf("String dose'nt have double or triple String !!!")
	}
}