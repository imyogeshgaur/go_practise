package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pointerOfUser := bufio.NewReader(os.Stdin);
	pointerOfChoice := bufio.NewReader(os.Stdin);
	pointerOfNumber1 := bufio.NewReader(os.Stdin);
	pointerOfNumber2 := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter Your Name : ");

	nameOfUser, _ := pointerOfUser.ReadString('\n');

	fmt.Println("Welcome to go Calculator !!!",nameOfUser)

	
	fmt.Printf("Enter Your First Number : \n");
	pointerOfNumber1Str ,_ := pointerOfNumber1.ReadString('\n');

	number1OfUser, err := strconv.ParseInt(strings.TrimSpace(pointerOfNumber1Str),10,64);

	if err != nil{
		fmt.Println("Error in entering the first number !!!")
		fmt.Println(err);
	}else{
		fmt.Println("Your first Number : ",number1OfUser);
	}

	fmt.Printf("Enter Your Second Number : \n");
	pointerOfNumber2Str ,_ := pointerOfNumber2.ReadString('\n');

	number2OfUser, err := strconv.ParseInt(strings.TrimSpace(pointerOfNumber2Str),10,64);

	if err != nil{
		fmt.Println("Error in entering the seconf number !!!")
		fmt.Println(err);
	}else{
		fmt.Println("Your Second Number : ",number2OfUser);
	}


	fmt.Println("\nEnter 1 For Addition !!!")
	fmt.Println("Enter 2 For Substraction !!!")
	fmt.Println("Enter 3 For Multiplication !!!")
	fmt.Println("Enter 4 For Division !!!")
	fmt.Println("Enter 5 For Modulous !!!")

	fmt.Printf("Enter Your Choice : ");

	choiceOfUserStr ,_ := pointerOfChoice.ReadString('\n');

	choiceOfUser, err := strconv.ParseInt(strings.TrimSpace(choiceOfUserStr),10,64);

	if err != nil{
		fmt.Println("Error in entering the choice !!!")
		fmt.Println(err);
	}else{
		switch choiceOfUser {
		case 1: fmt.Printf("You choose addition and the anwer is %d",number1OfUser+number2OfUser)
				break;
		case 2: fmt.Printf("Tou choose substraction and the anwer is %d",number1OfUser-number2OfUser)
				break;
		case 3: fmt.Printf("You choose multiplication and the anwer is %d",number1OfUser * number2OfUser)
				break;
		case 4: fmt.Printf("You choose division and the anwer is %d",number1OfUser / number2OfUser)
				break;
		case 5: fmt.Printf("You choose modulo and the anwer is %d",number1OfUser % number2OfUser)
				break;
		default: fmt.Printf("Sorry Wrong Input :( ");
				os.Exit(0)
		}
	}

	
	




}