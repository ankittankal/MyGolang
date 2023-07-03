package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	welcome := "Hello welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of product :")

	//comma ok syntax / err ok syntax

	input ,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, " , input)
	fmt.Printf("Type of the rating is: %T" ,input)

	numRating , err := strconv.ParseFloat(strings.TrimSpace(input),64) 

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("Added 1 to your rating" , numRating + 1)
	}
}