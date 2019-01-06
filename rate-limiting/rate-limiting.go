// rate limiting is an important mechanism for controlling resource utilization
// and maintaining quality of service.
// Go supports rate limiting with goroutines, channels and tickers.

package main

import "time"
import "fmt"

func main() {

	// limit handling of incoming requests.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// this limiter channel will receive a value every 200 milliseconds.
	// this is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on a received from the limiter channel b efore serving
	// eaching request, we limit ourselves to 1 request every 200 ms.
	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}

	// we may want to allow short bursts of requests in our rate limiting
	// scheme while preserving the overall rate limit. we can accomplish this
	// by buffering our limiter channel. this busrstyLimiter channnel will
	// allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200 ms we will incoming requests. the first 3 of these will
	// benifit from the burst capability of bustyLimiter.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
