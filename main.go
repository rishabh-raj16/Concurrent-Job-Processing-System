package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id     int
	Status string
}

func main() {
	// learn importance of channel
	jobs := make(chan Job, 100)
	var wg sync.WaitGroup
	for i:=1;i<3;i++{
		// when u put go key words  before any task make it goroutine 
		// which executes asynchronously (in the background) 
		// start worker
		
		go worker(i,jobs,&wg)
		
		
	}
	// adding jobs to channel
		for i:=1;i<20;i++{
			wg.Add(1)
			jobs <- Job{
				Id: i,
				Status: "Queued",
			}
		}

		close(jobs) // close job after all job is sent

		// keep main alive temporarily
		wg.Wait()
		fmt.Println("All jobs completed")

}

// what is this <-chan
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	
	// for job := range jobs 
	for
	{
		job,ok:=<- jobs
		fmt.Println("job",job,"ok",ok)

		if !ok{
			fmt.Println("worker",id,"channel closed")
			return
		}
		
		// lean fmt vs log package
		fmt.Println("worker", id, "processing", job.Id)
		// this time.sleep is to wait for two sec so it does not finish immediately
		  time.Sleep(2 * time.Second)
		fmt.Println("worker",id,"finished",job.Id)
		wg.Done()
	}
}