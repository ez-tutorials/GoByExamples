package main

import "fmt"

/*
	"for" is Go's only looping construct.
	3 types of for loops
*/

func main() {
	i := 1
	//01. Basic type: single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//02, Classic: initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// for without a condition will loop repeatedly until you break out of
	// the loop or return from the enclosed function.
	count := 0
	for {
		fmt.Println("Loop: ", count)
		count++
		if count == 10 {
			break
		}
	}
	// continue to next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
