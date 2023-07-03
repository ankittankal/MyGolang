package main

import "fmt"

func doesPetExists(pets map[string]string, pet string) bool {
	_, ok := pets[pet]

	return ok
}

func main() {
	pets := make(map[string]string)

	pets["rambo"] = "dog"
	pets["sairon"] = "cat"

	fmt.Println(doesPetExists(pets, "sairon"))
}
