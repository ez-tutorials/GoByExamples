package main

import "fmt"

/*
	"for" is Go's only looping construct.

*/


func main() {
	i := 1
	// Basic type: single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic: initial/condition/after for loop
	for j:=7; j<=9; j++ {
		fmt.Println(j)
	}
	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosed function.
	for {
		fmt.Println("Loop")
		break
	}
	// continue to next iteration
	for n:=0; n<=5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
