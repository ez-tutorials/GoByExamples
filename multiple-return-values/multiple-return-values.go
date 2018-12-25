// builtin support for multiple return values

package main

import "fmt"

// (int, int) in the function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// if you only want a subset of the returned values, use the blank
	// identifier
	_, c := vals()
	fmt.Println(c)
}
