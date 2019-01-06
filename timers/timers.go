// timers are for when you want to do something once in the future.
package main

import (
	"fmt"
	"time"
)

func main() {

	/* Timer represent a single event in the future.
	* It needs to be told how long it must wait, and it provides a channel that will be notified at that time.
	*  */
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer's channel C until it sends a value
	// indicating that the time expired.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// time.Sleep() waits a given number of second
	// One reason a timer may be useful is that it can be cancelled the timer
	// before it expires.

	// The first timer will expire ~2s after we start the program,
	// but the second should be stopped before it has a chance to expire.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
