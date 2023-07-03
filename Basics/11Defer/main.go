package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Hello")
	fun1()
}

//defer follows LIFO order

func fun1()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
