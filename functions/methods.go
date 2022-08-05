package main

import (
	"fmt"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func main() {
	fnMethods()
}

type stringy func(string) string

func (s stringy) printString(str string) {
	fmt.Println(s(str))
}

func fnMethods() {
	var myStringy stringy = func(str string) string {
		return "Hello " + str
	}

	// s := myStringy("Gophers")
	// fmt.Println(s)

	myStringy.printString("Gophers")
}
