// channels can be used to synchronize execution across goroutines.
package main

import "fmt"
import "time"

// this func will run in a goroutine
// the done channel will be used to notify another goroutine that this
// function's done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// send a value to notify that it's done
	done <- true
}

func main() {
	done := make(chan bool, 1)
	// start a worker goroutine, give it the channel to notify on
	go worker(done)
	// block until we receive a notification from the worker on the channel.
	<-done
	// if the avbove line is removed, the progrom will exit before the worker
	// event.
}
