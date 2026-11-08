package main

import "fmt"

// var someName := "Sarah" THIS is wrong we cannot use := to declare variable out of teh function
func main() {
	//Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	//Add new Line automatically
	fmt.Println("Hello ninjas!")
	fmt.Println("Goodbye ninjas!")

	// print variables
	age := 20
	name := "shone"
	println("my age is", age, "and my name is", name)

	// Printf() Formating string %v default specifier
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	// %q will not work for age since it works for prints a string surrounded by quotation marks
	fmt.Printf("my age is %q and my name is %q\n", age, name)

	// Print the type of veriable
	fmt.Printf("Age of type %T\n", age)

	// Print floats
	fmt.Printf("you scored %f points \n", 225.55)
	fmt.Printf("you scored %0.1f points \n", 225.55)

	// Sprinttf (save formated strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("The saved string is:", str)
}
