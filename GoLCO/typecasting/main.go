package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//User Input 

	inputPointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter Your name : ");

	input, _ := inputPointer.ReadString('\n');

	fmt.Printf("Your name is : %s",input)

	//Typecasting is necessary because the terminal is using string as default type

	//String to int64
	inputVal := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter The Number : ");

	inputStr , _  := inputVal.ReadString('\n');

	inutNum , err := strconv.ParseInt(strings.TrimSpace(inputStr),10,64);

	if err != nil {
		fmt.Print(err)
	}else{
		fmt.Println("The number is : ",inutNum);
	}

	//String to float
	inputValFloat := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter The Number : ");

	inputStrFloat , _  := inputValFloat.ReadString('\n');

	inutNumFloat , err := strconv.ParseFloat(strings.TrimSpace(inputStrFloat),64);

	if err != nil {
		fmt.Print(err)
	}else{
		fmt.Println("The number is : ",inutNumFloat);
	}
	
}