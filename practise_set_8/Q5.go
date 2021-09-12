package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TommyVecetti struct {
	characterState string
}

func setHiting(p TommyVecetti) {
	p.characterState = "Hiting"
	fmt.Printf("The mode of the Character is %v", p.characterState)
}
func setRunning(p TommyVecetti) {
	p.characterState = "Runing"
	fmt.Printf("The mode of the Character is %v", p.characterState)
}
func setFiring(p TommyVecetti) {
	p.characterState = "Firing"
	fmt.Printf("The mode of the Character is %v", p.characterState)
}

func main() {
	ph := TommyVecetti{}

	choicePointer := bufio.NewReader(os.Stdin)

	fmt.Println("Press 1 for Hiting mode !!!")
	fmt.Println("Press 2 for Runing mode !!!")
	fmt.Println("Press any other number for Firing mode !!!")
	fmt.Printf("Enter the State of Character : ")

	choiceReader ,_ := choicePointer.ReadString('\n');

	choice ,err := strconv.ParseInt(strings.TrimSpace(choiceReader),10,64);

	if err == nil {
		if(choice==1){
			setHiting(ph)
		}else if(choice==2){
			setRunning(ph)
		}else{
			setFiring(ph)
		}
	}
}