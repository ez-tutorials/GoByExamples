// Anonymous functions
// forms closures
// used to define a function inline without having to name it. (lambda)

package main

import "fmt"

// this function returns another function
// the returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// we assign the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each
	// time nextInt is called.
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}
