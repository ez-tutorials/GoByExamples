// tickers are for when you want to do something repeatedly at regular
// intervals.

package main

import "time"
import "fmt"

func main() {
	// tickers use a similar mechanism to timer: a channel that is sent values.
	// We use the range builtin on the channel to iterate over the values as
	// they arrive very 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	// tickers can be stopped.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
