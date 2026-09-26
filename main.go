package main

import (
	"fmt"
	"time"
)

type Job struct {
	Id     int
	Status string
}

func main() {
	// learn importance of channel
	jobs := make(chan Job, 100)
	for i:=1;i<3;i++{
		// when u put go key words  before any task make it goroutine 
		// which executes asynchronously (in the background) 
		// start worker
		go worker(i,jobs)
		
	}
	// adding jobs to channel
		for i:=1;i<20;i++{
			jobs <- Job{
				Id: i,
				Status: "Queued",
			}
		}

		// keep main alive temporarily
		time.Sleep(15 * time.Second)

}

// what is this <-chan
func worker(id int, jobs <-chan Job) {
	
	for job := range jobs {
		// lean fmt vs log package
		fmt.Println("worker", id, "processing", job.Id)
		// this time.sleep is to wait for two sec so it does not finish immediately
		time.Sleep(2 * time.Second)
		fmt.Println("worker",id,"finished",job.Id)
	}
}