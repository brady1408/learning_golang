package main

import "fmt"

func main() {
	printSum("First List", 1, 2, 3)
	list := []int{1, 2, 3, 4}
	printSum("Second Slice", list...)
}

func printSum(s string, i ...int) {
	var sum int
	for _, v := range i {
		sum += v
	}

	fmt.Printf("the sum of %s is %d\n", s, sum)
}
