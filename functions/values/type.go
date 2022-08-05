package main

import (
	"fmt"
	"reflect"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func main() {
	impl()
	expl()
}

func impl() {
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3))
	// f = product(3, 3)
	// fmt.Println(f(3, 3))
}

func expl() {
	type fn func(int) int
	//var f fn
	var f func(int) int
	fmt.Printf("%v\n", reflect.TypeOf(f))
	fmt.Printf("%v\n", reflect.TypeOf(square))
	fmt.Printf("%v\n", reflect.TypeOf(negative))
	fmt.Printf("%v\n", reflect.TypeOf(product))
}
