package main

import "fmt"

// typical function signature
func typical(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("add", add(10, 5))
	fmt.Println("sub", sub(10, 5))
	fmt.Println("sub2", sub2(10, 5))
	fmt.Println("first", first(10, 5))
	fmt.Println("zero", zero(10, 5))
	// Cannot specify arguments by name, this is not valid go syntax
	// fmt.Println("defaultValues", defaultValues(b = 5, a = 10))
	a, b := 10, 5
	fmt.Println("before reverse", a, b)
	reverse(a, b)
	fmt.Println("after reverse", a, b)
	realReverse(&a, &b)
	fmt.Println("after realReverse", a, b)
	noReturn()
}

// More Signatures
func a(i, j, k, int, s, t string) { fmt.Print("This signature combines parameters with like types") }
func add(x int, y int) int        { return x + y }

// Return values can be named just like parameters
func sub(x int, y int) (z int)  { z = x - y; return }
func sub2(x int, y int) (z int) { z = x - y; return -1 }

// What are the use cases for this? Maybe matching a signature type fSig func(int, int) int.
func first(x int, _ int) int { return x }

// Parameter names are not required only the type, you cannot reference a parameter that is missing a name.
// Satisfy interface? Satisfy type?
func zero(int, int) int { return 0 }

// This is not a legal function declaration
// func defaultValues(a int = 1, b int = 2) int { return 0}

// parameters are local variables
func reverse(a int, b int) { a, b = b, a; fmt.Println("reverse", a, b) }

// reference parameters used by caller are can be affected by the function
func realReverse(a *int, b *int) { *a, *b = *b, *a; fmt.Println("realReverse", *a, *b) }

// a function does not need a return if the code will never exit
func noReturn() int { panic(0) }
