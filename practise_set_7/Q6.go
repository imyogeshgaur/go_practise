package main

import "fmt"

func getAverageOfNumber(arr [5] float64) float64 {
	var sum float64= 0;
	for i := 0; i < 5; i++ {
		sum += arr[i];
	}
	var avg float64 = sum /5;
	return avg

}

func main() {
    var array = [5] float64{1,2,3,4,5};
	averageOfArray := getAverageOfNumber(array)

	fmt.Printf("The Average of Array is %f", averageOfArray)
	
}