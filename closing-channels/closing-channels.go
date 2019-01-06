// closing a channel indicates that no more values will be sent on it.
// this can be useful to communicate completion to the channel's receivers.

package main

import "fmt"

func main() {
	// in this example we'll use a jobs channel to communicate work to be done
	// from the main() goroutine to a worker goroutine.
	// when we have no more jobs for the worker we'll close the jobs channel.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// this is the worker goroutine.
	// it repeatedly receiveds from jobs with j, more := <-jobs.
	// in this special 2-value form of receive, the mroe value will be false if
	// jobs has been closed and all values in the channel have already been
	// received. we use this to notify on done when we have worked all our
	// jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// this sends 3 jobs to the worker over the jobs channel, then close it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	// wait the worker using the synchronization approach.
	<-done
}
