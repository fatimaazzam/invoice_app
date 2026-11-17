package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye, %s!\n", name)
}

func cycleName(names []string, f func(string2 string)) {
	for _, name := range names {
		f(name)
	}
}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleName([]string{"cloud", "tifa", "baret"}, sayGreeting)
	cycleName([]string{"cloud", "tifa", "baret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Printf("Circle Area: %.3f\n", a1)
	fmt.Printf("Circle Area: %.3f\n", a2)
}
