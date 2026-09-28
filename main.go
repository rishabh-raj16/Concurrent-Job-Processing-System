package main

import (
	"fmt"
	"sync"
	// "time"
)

type Job struct {
	Id     int
	Status string
}

func main() {
	jobs := make(chan Job, 100)
	completedJob:=0

	var wg sync.WaitGroup
	var mu sync.Mutex
	for i:=1;i<3;i++{
		// when u put go key words  before any task make it goroutine 
		// which executes asynchronously (in the background) 
		// start worker
		
		go worker(i,jobs,&wg, &completedJob, &mu)
		
		
	}
	// adding jobs to channel
		for i:=1;i<100;i++{
			wg.Add(1)
			jobs <- Job{
				Id: i,
				Status: "Queued",
			}
		}

		close(jobs) // close job after all job is sent

		// keep main alive temporarily
		wg.Wait()
		// fmt.Println("All jobs completed")
		fmt.Println("Completed jobs:", completedJob)

}

// what is this <-chan
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup , completedJob *int, mu *sync.Mutex) {
	
	for job := range jobs {
	// for
	// {
	// 	job,ok:=<- jobs
	// 	fmt.Println("job",job,"ok",ok)

	// 	if !ok{
	// 		fmt.Println("worker",id,"channel closed")
	// 		return
	// 	}
		
		fmt.Println("worker", id, "processing", job.Id)
		// this time.sleep is to wait for two sec so it does not finish immediately
		//   time.Sleep(1 * time.Second)
		fmt.Println("worker",id,"finished",job.Id)

		mu.Lock()

		(*completedJob)++
		mu.Unlock()
		wg.Done()
	}
}