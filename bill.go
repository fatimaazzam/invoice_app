package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b

}

// RECEIVE FUNCTION
// Inorder to call the function with specific type we use receive function as follows
// Format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) // -25 it adds spaces at the right just for alignments
		total += v
	}
	fs += fmt.Sprintf("%-25v ... $%0.2f", "Total:", total) //-25 it adds spaces at the right just for alignments
	return fs
}
