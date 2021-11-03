package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, hasMore := <-jobs
			time.Sleep(1 * time.Second)
			if hasMore {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		fmt.Println("sending job", j)
		time.Sleep(1 * time.Second)
		jobs <- j
		time.Sleep(1 * time.Second)
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
