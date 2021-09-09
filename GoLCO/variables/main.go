package main

import "fmt";

func main()  {

	//Strings 

	var nameOfUser1 = "Yogesh";  //1st wayof writing the variables 
	var nameOfUser2 string = "Ramesh"; //2nd Way of writing the variables 
    nameOfUser3 := "Dinesh"; //3rd way of writing the variables 

	fmt.Printf("The first user name is %s \nThe Second user name is %s \nThe Third user name is %s \nThe type of variables is %T\n",nameOfUser1,nameOfUser2,nameOfUser3,nameOfUser1);

	//Unsigned Integers are of following types 

	 var age1 uint8 = 12; 
	 var age2 uint16 = 1267;
	 var age3 uint32 = 1276342;
	 var age4 uint64 = 1273465547654765453;
	 age := 0;

	 fmt.Printf("The first user age is %d \nThe Second user age is %d \nThe Third user age is %d \nThe fourth user age is %d \nThe type of variables is %T\n",age1,age2,age3,age4,age);
	
	 //Signed Integers are of following types 

	 var age6 int8 = -12;
	 var age7 int16 = -1267;
	 var age8 int32 = -1276342;
	 var age9 int64 = -1273465547654765453;
	 age10 := -1;

	 fmt.Printf("The first user age is %d \nThe Second user age is %d \nThe Third user age is %d \nThe fourth user age is %d \nThe type of variables is %T\n",age6,age7,age8,age9,age10);
	
	 // Floating Point Numbers 

	 var salary1 float32 = 53485.654353675356354;
	 var salary2 float64 = 53485.654353675356354;
	 salary := 0; 
	
	 fmt.Printf("The first user age is %f \nThe Second user age is %f \nThe type of variables is %T",salary1,salary2,salary);

	//Complex Numbers 

	 var number1 complex64 = 34+45.6789033535i;
	 var number2 complex128 = 34+45.6789033535i;

	 fmt.Println("\nThe First number is ",number1)
	 fmt.Println("The second number is ",number2)

	 //Default values (if not initialzed)

	 var mystring string;  // " "
	 var mycomplex128 complex128; //0 + 0i
	 var mycomplex64 complex64; //0 + 0i
	 var myuint8 uint8; // 0
	 var myuint16 uint16; // 0
	 var myuint32 uint32; // 0
	 var myuint64 uint64; // 0
	 var myint8 int8; // 0
	 var myint16 int16; // 0
	 var myint32 int32; // 0
	 var myint64 int64; // 0
	 var mybool bool; // false
	 var myfloat32 float32; // 0
	 var myfloat64 float64; // 0

	//  fmt.Println("\nThe default value of string is ",mystring);
	//  fmt.Println("The default value of complex128 is ",mycomplex128);
	//  fmt.Println("The default value of complex64 is ",mycomplex64);
	//  fmt.Println("The default value of uint8 is ",myuint8);
	//  fmt.Println("The default value of uint16 is ",myuint16);
	//  fmt.Println("The default value of unint32 is ",myuint32);
	//  fmt.Println("The default value of unint64 is ",myuint64);
	//  fmt.Println("The default value of int8 is ",myint8);
	//  fmt.Println("The default value of int16 is ",myint16);
	//  fmt.Println("The default value of int32 is ",myint32);
	//  fmt.Println("The default value of int64 is ",myint64);
	//  fmt.Println("The default value of float32 is ",myfloat32);
	//  fmt.Println("The default value of float64 is ",myfloat64);
	//  fmt.Println("The default value of bool is ",mybool);

}