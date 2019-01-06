package main

import "fmt"
import "time"

// we'll run serveral concurrent instances of this function.
// this workers will received work on the jobs channel and send the
// corresponding results on results.
// it will sleep 1 second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// we need to send them work and collect their results.
	// we made 2 channels
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// this starts up 3 workers, initially blocked because there are no job
	// yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// here we send 5 jobs and then close channel to indicate that's all work
	// we have.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// colect all results of the work.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
