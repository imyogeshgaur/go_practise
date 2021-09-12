package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Phone struct {
	phoneState string
}

func setRinging(p Phone) {
	p.phoneState = "Ringing"
	fmt.Printf("The mode of the Phone is %v", p.phoneState)
}
func setVibrating(p Phone) {
	p.phoneState = "Vibrating"
	fmt.Printf("The mode of the Phone is %v", p.phoneState)
}
func setSilent(p Phone) {
	p.phoneState = "Silent"
	fmt.Printf("The mode of the Phone is %v", p.phoneState)
}

func main() {
	ph := Phone{}

	choicePointer := bufio.NewReader(os.Stdin)

	fmt.Println("Press 1 for Ringing mode !!!")
	fmt.Println("Press 2 for Vibrating mode !!!")
	fmt.Println("Press any other number for Silent mode !!!")
	fmt.Printf("Enter the State of Phone : ")

	choiceReader ,_ := choicePointer.ReadString('\n');

	choice ,err := strconv.ParseInt(strings.TrimSpace(choiceReader),10,64);

	if err == nil {
		if(choice==1){
			setRinging(ph)
		}else if(choice==2){
			setVibrating(ph)
		}else{
			setSilent(ph)
		}
	}
}