package main

import (
	"fmt"
)

func main() {
	x := 0

	// In GO we use for as while
	fmt.Println("Using For as While")
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	// Using For loop
	fmt.Println("Using For loop")
	for i := 0; i < 5; i++ {
		fmt.Println("value of x is: ", i)
	}

	//Loop through slice of strings
	fmt.Println("Loop through slice")
	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//Using Loop with range
	fmt.Println("Using range with loop")
	for index, value := range names {
		fmt.Printf("The value at index %v  is %v \n", index, value)

	}

	// remove the index
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		value = "new String" // This will not update the original value in the slice because value is a local variable
	}

	fmt.Println(names)

}
