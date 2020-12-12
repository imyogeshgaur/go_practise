package main

import (
	"fmt"
	"strings"
)

func main(){
   var userString,lowerCaseString string
   fmt.Printf("Enter the Sting : ")
   fmt.Scanln(&userString)
   lowerCaseString  = strings.ToLower(userString)
   fmt.Printf("The Lower Case String is : %s ",lowerCaseString)
}