package main

import "fmt"

func main()  {
	
	//Initialisation
	var fruits [4]string
	
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[3] = "peach"

	fmt.Println(fruits)
	//prints the size of array
	fmt.Println(len(fruits))

}