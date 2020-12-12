package main

import "fmt"

func main(){
	var originalGrade,encryptGrade int
	fmt.Printf("Enter the Marks of the Student :")
	fmt.Scanln(&originalGrade)
	encryptGrade = originalGrade + 8
	fmt.Printf("The Encrypted Marks are ",encryptGrade)
	fmt.Printf("The Decrypted Marks are ",originalGrade)
}
