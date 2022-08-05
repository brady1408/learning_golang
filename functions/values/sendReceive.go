package main

import (
	"fmt"
)

func hello(s string) string {
	return "Hello " + s
}

func main() {
	fmt.Println(takeFunc(hello))
	fmt.Println(returnFunc()("from return fuction"))
}

func takeFunc(f func(string) string) string {
	return f("from take function")
}

func returnFunc() func(string) string {
	return hello
}
