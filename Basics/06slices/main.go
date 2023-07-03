package main

import (
	"fmt"
	"sort"
)


func main()  {
	
	//initialisation
	var fruits = []string{"alpples" , "tomato" , "peach"}
	fmt.Println(fruits)

	//append
	fruits = append(fruits, "Banana")
	fmt.Println("aftr appending :",fruits)

	//slice
	fruits = fruits[1:]
	fmt.Println("after slicing :",fruits)


	//slice iinteresting property
	scores := make([]int , 4)

	scores[0] = 10
	scores[1] = 20
	scores[2] = 30
	scores[3] = 40
	//This will throw an error because memory only allcated for 4 values
	//scores[4] = 50

	//But when we append it will reallocate the memory 
	scores = append(scores,55,66,77)
	fmt.Println(scores)

	//sort 
	sort.Ints(scores)
	fmt.Println(scores)

 }