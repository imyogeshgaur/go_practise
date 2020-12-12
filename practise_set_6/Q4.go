package main

import "fmt"

func main(){
	var mat1[2][3],mat2[2][3],ans[2][3] int
	
	fmt.Print("Enter the First Matrix : ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Enter the Value at (",i",",j")")
			fmt.Scanln(&mat1[i][j])
		}
	}
	
	fmt.Print("Enter the Second Matrix : ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Enter the Value at (",i",",j")")
			fmt.Scanln(&mat2[i][j])
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			mat3[i][j] = mat1[i][j] +mat2[i][j]
		}
	}
	fmt.Print("The Resultant matrix is ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(mat3[i][j]," ")
		}
		fmt.Println(" ")
	}
}