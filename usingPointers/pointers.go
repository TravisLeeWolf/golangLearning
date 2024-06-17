package main

import "fmt"

func main() {
	var age int
	fmt.Println("How old are you?: ")
	fmt.Scan(&age) // Were using the pointer of age to override the value

	agePointer := &age

	fmt.Println("You're:", *agePointer)
	getYearsAsAdult(agePointer)
	fmt.Println("You've been an adult for:", age)
}

func getYearsAsAdult(age *int) {
	*age = *age - 18 // Overrides age
	// return *age - 18
}
