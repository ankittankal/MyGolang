package main

import "fmt"

func average(num1 float32, num2 float32, num3 float32) float32 {
	var avg = (num1 + num2 + num3) / 3
	return avg
}

func main() {
	//Ankit := User{"Ankit", true, 5}
	//GetStatus(Ankit)
	fmt.Println(average(5, 5, 100))
}

// type User struct {
// 	Name   string
// 	Status bool
// 	Age    int
// }

// func GetStatus(u User) {
// 	fmt.Println(u.Status)
// }
