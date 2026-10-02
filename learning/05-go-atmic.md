## Memory Primitive

Atomic operations allow multiple goroutines to safely modify a shared variable without explicitly using a mutex.

Atomic operations are supported using low-level CPU/hardware synchronization instructions. An atomic read-modify-write operation behaves as one indivisible operation from the perspective of other goroutines.

For example, an ordinary increment conceptually works like:

`read → modify → write`

Another goroutine can access the same variable between these steps, which can create a data race and lost updates.

But an atomic operation like `Add(1)` is indivisible, so another atomic operation on the same value cannot intervene in the middle of that read-modify-write operation.

Even if two CPU cores try to perform the atomic operation at almost exactly the same time, the hardware ensures that their modifications are ordered correctly.

## Use

Atomic operations are best for simple shared state such as:

- counters
- flags
- simple numeric values

For example:

`completedJobs.Add(1)`

If multiple variables must stay consistent **together**, making each variable atomic does not automatically make the whole operation atomic.

For example, if `balance` and `transactionCount` must be updated as one consistent operation, separate atomic operations can still allow another goroutine to observe an intermediate state.

In such cases, it is better to use `sync.Mutex` / `sync.RWMutex` to protect the complete shared state or critical section.