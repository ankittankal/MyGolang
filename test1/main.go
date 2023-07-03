package main

import "fmt"

func calculateAvg(arr1 [5]int) int {
	sum := 0
	for _, num := range arr1 {
		sum += num
	}
	return sum
}

func main() {

	//var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	//var arr1 = [5]int{1, 2, 3, 4, 5}
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(calculateAvg(arr1))

	fmt.Println(arr1)
}
