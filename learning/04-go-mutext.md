## mutex
## Mutex

1. Mutex is used when we want to restrict multiple goroutines from accessing a shared resource or critical section at the same time. Only one goroutine should access that protected section at a time.

2. Mutex does not itself lock a variable or resource. The mutex itself becomes locked/occupied when a goroutine calls `Lock()`. If another goroutine reaches `Lock()` on the same mutex, it has to wait until the mutex is unlocked.

3. The goroutine that locks the mutex should unlock it after completing the protected work. If the mutex is never unlocked, other goroutines waiting for that mutex can remain blocked and this can contribute to a deadlock.

4. Mutex is part of Go's `sync` package.

Example:

```go
mu.Lock()
(*completedJob)++
mu.Unlock()
```

Here `completedJob` itself is not locked. We use `mu` to make sure only one goroutine at a time executes the protected update.

## sync.RWMutex
RWMutex give us two king of locking 
if two worker wants to read a resource it allow 
but if any two workers wants to write at a same time or one wants to read and another wants to write 
it does not allow at the same time , bcz without synchronization its a data race in go

mu.RLock() // read lock
mu.RUnlock()

mu.Lock()
mu.Unlock() // write lock

## scheduling order
we know that RWMutex allow all the read together , so when a wite comes and call mu.Lock()
mutex check if any goroutine is reading then it will make write to wait and also if another read comes even though it does not block reads when another worker is reading but this time it will make read to wait , when the first read mu.RUnlock() then write will acquire mutex and all the read will wait 

RWmutex does not give guarantee to execute worker in the order it came
means if this is arrival order 
R1 → W1 → R2,R3 → W2 → R4,R5
its does not guarantee to be executed and same order
 it only guarantee that 
 if readers that Already have RLock() can finish its job
 if new reader comes he can have access of RLock() without waiting
 but
 when a write comes and goes in waiting state no new reader can access the RLock()
until wite unlock it .
after this cycle again  controls ACCESS will be empty and any worker can acquire in any order 
means if  read acquire RLock() and write went to waiting stage then other worker have to wait to complete this cycle 
same if one write acquire first the all other read/write have to wait to finish the write its job .
