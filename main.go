package main

import "fmt"

// var someName := "Sarah" THIS is wrong we cannot use := to declare variable out of teh function
func main() {

	//Strings
	var nameOne string = "mario"
	var nameTwo = "John"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Nohu"
	nameThree = "Lona"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Moses"
	fmt.Println(nameFour) // Cannot be used out of teh function

	// integers
	var ageOne int = 20
	var ageTwo int = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

}
