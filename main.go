package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	// Group A types -> strings, ints, bools, floats, array, struct (NON-POINTERS VALUES)
	name := "tifa"

	//updateName(name)
	fmt.Println("memory address of the name is:", &name)

	m := &name
	fmt.Println("memory address of the name is:", m)
	fmt.Println("memory address of the name is:", *m)
	fmt.Println("memory address of the m is:", &m)

	fmt.Println("Before Update: ", name)
	updateName(m)
	fmt.Println("After Update: ", name)

}
