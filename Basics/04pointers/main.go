package main

import "fmt"

func main()  {
	fmt.Println("Welcome to pointers")

	//initialising variable
	//one := 2

	//initialising pointer
	var ptr *int 
	fmt.Println("Value of pointer is :",ptr)

	//playing with pointers
	myNum := 25
	var ptr2 = &myNum
	fmt.Println("value of pointer address is :",ptr2)
	fmt.Println("value of pointer address is :",*ptr2)

}