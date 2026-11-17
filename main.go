package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	namesSlice := strings.Split(s, " ")

	var initials []string
	for _, v := range namesSlice {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
func main() {
	fn1, sn1 := getInitials("tifs lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud lockhart")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("luji")
	fmt.Println(fn3, sn3)

}
