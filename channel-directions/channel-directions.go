// when using channels as function parameters, a channel can be specified to
// receive or send values.

package main

import "fmt"

// this ping function only accepts a channel for sending values.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// the pong function accepts one channel for receives (pings) and a second for
// sends (pong)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
