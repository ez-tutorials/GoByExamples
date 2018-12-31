// timeouts are important for programms that connect to external resources or
// that otherwise need to bound execution time.

package main

import "time"
import "fmt"

func main() {
	// we're executing an external call that returns its result on a channel c1
	// after 2s
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	// select implementing a timeout
	select {
	case res := <-c1: // waits the result
		fmt.Println(res)
	case <-time.After(1 * time.Second): // waits a value to be sent after the timeout of 1s
		fmt.Println("timeout 1")
	}

	// if we allow a longer timeout of 3s, then the receive from c2 will
	// succeed and we'll print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
	// using this select thimeout pattern requires communicating results over
	// channels.
	//
}
