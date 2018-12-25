package main

import "fmt"

// function that takes 2 ints and returns their sum
func plus(a int, b int) int {
	return a + b
}

// Go requires explicit returns
// when there are consecutive parameters of the same type, type names for the
// like-typed can be omitted up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
