package main

import (
	"fmt"
	"sort"
	"strings"
)

// var someName := "Sarah" THIS is wrong we cannot use := to declare variable out of teh function
func main() {
	greeting := "Hello there friends!"
	fmt.Println(strings.Contains(greeting, "Hello"))

	// using replace does not change the original string
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	fmt.Println("original string value =", greeting)

	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// Sorting
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages) // it alter the original slice and sort it
	fmt.Println(ages)

	// Return the index of teh sorted slice
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	// If the number does not exist it return index not exist
	index2 := sort.SearchInts(ages, 90)
	fmt.Println(index2) // it will return 8 which is not exist in the previous slice

	// Sort Strings
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)                              // it alter the original slice and sort it
	fmt.Println(sort.SearchStrings(names, "mario")) // it will print the index

}
