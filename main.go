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
		go worker(i,jobs)
	}

}

// what is this <-chan
func worker(id int, jobs <-chan Job) {
	//  when we use for range we take two variable like a,b:=range list{} why here one
	for job := range jobs {
		// lean fmt vs log package
		fmt.Println("worker", id, "processing", job.Id)
		// this time.sleep is to wait for two sec but why are we adding it inside worker
		time.Sleep(2 * time.Second)
		fmt.Println("worker",id,"finished",job.Id)
	}
}