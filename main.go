package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}
func main() {
	// Group A types -> strings, ints, bools, floats, array, struct (NON-POINTERS VALUES)
	name := "tifa"
	name = updateName(name)

	fmt.Println(name)

	// Group B types -> slice, maps, functions (POINTERS WRAPPER VALUES)
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
