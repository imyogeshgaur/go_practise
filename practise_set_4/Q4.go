package main

import "fmt"

func main(){
	var choice int
	fmt.Printf("Enter the Choice of User : ")
	fmt.Scanln(&choice)
    switch choice {
	case 1:fmt.Printf("Today is Monday !!")
		   break
	case 2:fmt.Printf("Today is Tueday !!")
		   break
	case 3:fmt.Printf("Today is Wednesday !!")
		   break
	case 4:fmt.Printf("Today is Thrusday !!")
		   break
	case 5:fmt.Printf("Today is Friday !!")
		   break
	case 6:fmt.Printf("Today is Saturday !!")
		   break
    case 7:fmt.Printf("Today is Sunday !!")
		   break
    default : fmt.Printf("Wrong Input !!!")		
	}	
}