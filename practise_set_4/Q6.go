package main

import (
	"fmt"
	"strings"
)

func main(){
	var urlOfWebsite string
	fmt.Printf("Enter the Url : ")
	fmt.Scanln(&urlOfWebsite)
    if strings.Contains(urlOfWebsite,"com"){
		fmt.Printf("It is a Commercial Website !!!")
	}else if strings.Contains(urlOfWebsite,"org"){
		fmt.Printf("It is a Organisation's Website !!!")
	}else if strings.Contains(urlOfWebsite,"in"){
		fmt.Printf("It is an Indian Website !!!")
	}else{
		fmt.Printf("Undetectable Extension !!")
	}
}