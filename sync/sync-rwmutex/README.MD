# Making the Case for `sync.RWMutex`

## about://sync.RWMutex

> ```go
> type RWMutex struct {
> // contains filtered or unexported fields
> }
> ```
>
> A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an
> arbitrary number of readers or a single writer. The **zero value** for a  RWMutex
> is an unlocked mutex.
>
> A RWMutex **must not be copied** after first use.
>
> If a goroutine holds a RWMutex for reading and another goroutine might call
> Lock, no goroutine should expect to be able to acquire a read lock until the
> initial read lock is released. In particular, this prohibits recursive read
> locking. This is to ensure that the lock eventually becomes available; a
> blocked Lock call excludes new readers from acquiring the lock.

## Exercise

Find the bugs in the counter package.

