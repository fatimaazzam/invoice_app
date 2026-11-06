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

	// bits & memory
	// var numOne int8 = 215 THIS declaration is wrong we can not assign numOne
	// to 215 because we use int8 and this will become out of the range, instead we can use int16
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256
	fmt.Println(numOne, numTwo, numThree)

	// Floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 323123123123.2
	scoreThree := 1.5 // The default is float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
