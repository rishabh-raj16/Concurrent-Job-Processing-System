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