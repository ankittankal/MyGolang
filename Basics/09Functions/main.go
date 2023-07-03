package main

import "fmt"

func average(nums ...float64) float64 {
	sum := 0.0
	//var avg float64
	for _, num := range nums {
		sum = sum + num
	}
	return sum / float64(len(nums))
}

func main() {
	//greet()
	//fmt.Println(add(2,3))
	//fmt.Println(addMultpleVals(1,2,3,4,5))
	fmt.Println(average(5.0, 5.0, 10, 10))
}

//we are not allowed to write functions nsde anotherr function

func greet() {
	fmt.Println("Hello")
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func addMultpleVals(values ...int) int {
	total := 0

	//Internaly it stores all the values in slice
	for _, val := range values {
		total += val
	}

	return total
}
