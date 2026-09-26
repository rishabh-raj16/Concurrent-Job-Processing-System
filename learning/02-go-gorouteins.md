// goroutines
## what i thought initially
 Adding go keyword before a function will make it run background and the program will continue running while that goroutine is executing.

## what actually happened 
I wrote: 
jobs := make(chan Job, 100)
	for i:=1;i<3;i++{
		
		go worker(i,jobs)
	}  
but the program exited immediately 
## why?
bcz main() itself in a goroutine called main goroutine
add another go routine inside it does not force main to wait for it to get completed 
when main return go program terminate even if another goroutine still not finished 
## what i learned 
go schedule a function to work concurrently
it does not mean that it will wait for the function to get finish 
or
keep the parent function alive until it gets complete 
## solution
i need to make main() synchronize by adding sync.WaitGroup so that main can wait for other go routines to finish 
 
i will add nxt.