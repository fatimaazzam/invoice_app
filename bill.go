package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
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
	fs += fmt.Sprintf("%-25v ...$%v \n", "Tip:", b.tip)
	fs += fmt.Sprintf("%-25v ... $%0.2f", "Total:", total) //-25 it adds spaces at the right just for alignments
	return fs
}

// Update the bill
// Pointers is used since we want to original bill not the copy
func (b *bill) updateTip(tip float64) {
	b.tip = tip

}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Save Bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved bill:", b.name)
}
