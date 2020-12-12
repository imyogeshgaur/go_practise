package main

import "fmt"

func main(){
	var incomeOfUser,taxOfUser float64
	fmt.Printf("Enter The Income : ")
	fmt.Scanln(&incomeOfUser)
	if incomeOfUser < 250000{
         fmt.Printf("You did'nt need to pay Tax !!")
	  }else if incomeOfUser >= 250000 && incomeOfUser <= 500000{
		      taxOfUser = incomeOfUser/20
		      fmt.Printf("Your Tax to be paid is : %.2f",taxOfUser)
		  }else if incomeOfUser >= 500000 && incomeOfUser <=1000000{
			     taxOfUser = incomeOfUser/5
			     fmt.Printf("Your Tax to be paid is : %.2f",taxOfUser)
			  }else{
				    taxOfUser = incomeOfUser*0.3
		            fmt.Printf("Your Tax to be paid is : %.2f",taxOfUser)

	}
	
}