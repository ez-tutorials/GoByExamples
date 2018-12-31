// channels are unbuffered by default.
// they'll only accept sends (chan <-) if there is corresponding receive (<-
// chan) ready to receive the sent value.
// buffered channels accept a limited number of values without a corresponding
// receiver for those values.

package main

import "fmt"

func main() {
	// make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// because this channel is buffered, we can send these values into the
	// channel without a corresponding receiver.
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
