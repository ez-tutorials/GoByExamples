// a goroutine is a lightweight thread of execution.

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// this how a func is called synchronously
	f("direct")

	// go f(s) invokes the function in a goroutine, it then will execute
	// concurrently with the calling one
	go f("goroutine")

	// start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
