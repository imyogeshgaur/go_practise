package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func converToFarnehiet(cel float64) float64 {
    fr := (cel*1.8)+32;
	return fr;
}

func main() {
	celcisusPointer := bufio.NewReader(os.Stdin);

	fmt.Printf("Enter the Tempreature in Celcius : ");

	celcisusReader ,_ := celcisusPointer.ReadString('\n');

	celcisus , err := strconv.ParseFloat(strings.TrimSpace(celcisusReader),64)

	if err == nil {
		farhnhiet := converToFarnehiet(celcisus);
		fmt.Printf("The Tempreature in Farhnheit is %f",farhnhiet);
	}
}