package main

import (
	"fmt"
	"strings"
)

func main(){
	var originalString,replacedString string
   fmt.Printf("Enter the String !!")
   fmt.Scanln(&originalString)
   replacedString = strings.ReplaceAll(originalString," ","_")
   fmt.Printf("Replaced String is : %s",replacedString)
}