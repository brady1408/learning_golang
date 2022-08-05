package main

import (
	"fmt"
)

func hello(s string) string {
	return "Hello " + s
}

func main() {
	// This time we will declare the function literal when we call takeFunc
	fmt.Println(takeFunc(func(s string) string {
		return "Hello " + s
	}))

	// the function literal is declared directly as a return in returnFunc
	fmt.Println(returnFunc()("from return fuction"))
	// fmt.Println(returnFunc("from return fuction")())

	// an anonymous function can be called directly after it is declared
	defer func() {
		fmt.Println("DONE!")
	}()
}

func takeFunc(f func(string) string) string {
	return f("from take function")
}

func returnFunc() func(string) string {
	// declare function literal as part of the return
	return func(s string) string {
		return fmt.Sprintf("Hello %s", s)
	}
}

// func returnFunc(s string) func() string {
// 	// declare function literal as part of the return
// 	return func() string {
// 		return fmt.Sprintf("Hello %s", s)
// 	}
// }
