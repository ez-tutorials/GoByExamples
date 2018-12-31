// Channels are the pipes that connect concurrent goroutines .
// You can send values into channels from one goroutine to another and receive
// those values into another goroutine.

package main

import "fmt"

func main() {

	// create a new channel: make(chan val-type)
	// channels are typed by the values
	messages := make(chan string)

	// send a value into a channel using the channel <- syntax.
	//
	go func() { messages <- "ping" }()

	// The <- channel syntax receives a value from the channel
	msg := <-messages
	fmt.Println(msg)

	// by default sends and receives block until both the sender and receiver
	// are ready.
	// This allows us to wait at the end of our program for the message without
	// having to use any other synchronization.
}
