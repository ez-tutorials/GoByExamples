// Go's select lets you wait on multiple channel operations.
package main

import "time"
import "fmt"

func main() {

	// select across 2 channels
	c1 := make(chan string)
	c2 := make(chan string)
	// each channel will receive a value after some amount of time
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// use select to await both of these values simultaneously, printing each
	// one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// the total execution tiem is only ~2 seconds sine both the 1 and 2 second
	// Sleeps execute concurrently.
}
