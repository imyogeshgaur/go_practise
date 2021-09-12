package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type penFunction interface {
	write()
	refill()
}

type Pen struct{
	penRefilState int;
}

func write(p Pen) {
	fmt.Printf("Let's Write because the refil is fill upto %v!!!",p.penRefilState)
}

func refill(p Pen) {
	fmt.Printf("Let's Refill because the refil is fill upto %v!!!",p.penRefilState)
}

func main() {
	p := Pen{};
	
	penStatePointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter The Pen State in Percentage : ");

	penStateReader, _ := penStatePointer.ReadString('\n');

	penState, err := strconv.ParseInt(strings.TrimSpace(penStateReader),10,64);

	if err == nil {
		if(penState != 0){
			p.penRefilState = int(penState)
			write(p);
		}else{
			p.penRefilState = int(penState)
			refill(p);
		}
	}
}