package main

import "fmt"

func main() {
	myBill := newBill("mario's bill")
	myBill.updateTip(10)
	myBill.addItem("pie", 5.99)
	myBill.addItem("cake", 3.99)
	fmt.Println(myBill.format())

}
