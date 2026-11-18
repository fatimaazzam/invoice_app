package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "_", v)
	}

	// integers as key type
	phonebook := map[int]string{
		123456: "mario",
		543790: "luigi",
		987653: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[123456])

	// Update the value inside the map
	phonebook[987653] = "bowser"
	fmt.Println(phonebook)

	phonebook[543790] = "yoshi"
	fmt.Println(phonebook)
}
