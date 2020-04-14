# Exploring the (x/)sync Package - Leipzig Golang Meetup #10

Examples, references and snippets for the short (virtual) presentation at the [Leipzig Golang Meetup](https://www.meetup.com/de-DE/Leipzig-Golang/events/268785531/).

## Peeking into the sync package

> Package sync provides basic synchronization primitives such as mutual exclusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.
>
> Values containing the types defined in this package **should not be copied**.
> 
> Source: [pkg.go.dev](https://pkg.go.dev/sync?tab=doc#pkg-overview)


Types defined in this package:

- Cond
  - Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.
- Locker (interface)
  - A Locker represents an object that can be locked and unlocked.
- [Map](./sync/sync-map)
  - Map is like a Go map[interface{}]interface{} but is safe for concurrent use by multiple goroutines without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.
- Mutex
  - A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
- [Once](./sync/sync-once)
  - Once is an object that will perform exactly one action.
- Pool
  - A Pool is a set of temporary objects that may be individually saved and retrieved.
- [RWMutex](./sync/sync-rwmutex)
  - A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer. The zero value for a RWMutex is an unlocked mutex.
- WaitGroup
  - A WaitGroup waits for a collection of goroutines to finish.

> **Note:** excluding sub-package `"atomic"`.

## Peeking into the x/sync package

> Package x/sync provides Go concurrency primitives in addition to the ones provided by the language and “sync” and “sync/atomic” packages.
>
> Source: [pkg.go.dev](https://pkg.go.dev/mod/golang.org/x/sync)

Packages defined:

- [errgroup](./x-sync/errgroup)
  - Package errgroup provides synchronization, error propagation, and Context cancelation for groups of goroutines working on subtasks of a common task.
- semaphore	
  - Package semaphore provides a weighted semaphore implementation.
- [singleflight](./x-sync/singleflight)
  - Package singleflight provides a duplicate function call suppression mechanism.
- syncmap	
  - Package syncmap provides a concurrent map implementation. (prototype, **deprecated** with the addition to the standard lib in Go v1.9)

For a comprehensive coverage of this package see [this](https://rodaine.com/2018/08/x-files-sync-golang/) awesome blog post.
