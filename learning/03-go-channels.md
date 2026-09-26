## channnel
A channel is a medium through which goroutines can communicate and synchronize by sending and receiving values. We don't need to manually use a mutex for the channel's send/receive operation.
## cheating channel
type Job struct {
	Id     int
	Status string
}
jobs := make(chan Job, 100)
1. we are creating a channel which is of struct Job type and it can have buffer of 100 job at a time 
2. capacity =100 jobs waiting in buffer at one time 
Not maximum total jobs processed = 100
## sending a channel value
for i:=0;i<20;i++{
    jobs <- Job{
        Id: i
        Status: "Queued"
    }
     //so this <- Job{} is use to send value
}
## Receive  channel data
u use same <- symbol but now at left side there is variable to receive it 
ch <- "Hello" // sending text value to ch
msg := <-ch // receiving that in variable msg

## passing inside a function argument 
there are three way to pass channel in a func argument 
1.  worker(jobs chan Job) // bidirectional
here inside worker function u can do both send and receive value from a channel 
2. worker(jobs <-chan Job) // consume jobs
here chan is after <-
now in this worker u can only consume job , can not send value to channel
3. produceJobs(jobs chan<- Job)
here u can see chan is before <- 
in this function u can u can send more job to channel 

## using channel inside code 
for i:=1;i<3;i++{
		
		go worker(i,jobs)
		
	}
	// adding jobs to channel
		for i:=1;i<20;i++{
			jobs <- Job{
				Id: i,
				Status: "Queued",
			}
		}

here u can see that we wrote first consuming function before sending function 
bcz  Workers are started first. They don't wait for the producer loop to finish. As soon as a job becomes available, an available worker can receive it and start processing it while the producer continues sending more jobs


this is not a universal way to only add first consumer function and send function u can write 
// adding jobs to channel
		for i:=1;i<20;i++{
			jobs <- Job{
				Id: i,
				Status: "Queued",
			}
		}
first also but be cautious if the no of work is less then buffer of channel there is no issue first all job will added to buffer and then consumer function will start executing 
but if no of jobs get increase more then buffer then it will create a deadlock 
bcz u can not add more then buffer limit at a time and until and unless this for loop gets completed consumer will not start it work 

so understanding the requirement take decision
if u keep the consumer 1st then producer and adding work is slow then consuming then it will work 
but if consumer task is time taking and slow then sending to channel it will created backpressure