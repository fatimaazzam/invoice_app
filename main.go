package main

import "fmt"

// var someName := "Sarah" THIS is wrong we cannot use := to declare variable out of teh function
func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	ages := [3]int{20, 25, 30}
	var names = [4]string{"yoshi", "mario", "sarah", "bowser"}
	names[1] = "luigi"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slicing (use under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	// The following will return new slice
	//append(scores, 85)

	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
