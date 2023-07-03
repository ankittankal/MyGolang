package main

import (
	"fmt"
)

func main()  {
	
	courses := []string{"react","javascript","golang","ruby"}
	fmt.Println(courses)

	index := 2
	courses = append(courses[:index] , courses[index+1:]...)
	fmt.Println(courses)
}