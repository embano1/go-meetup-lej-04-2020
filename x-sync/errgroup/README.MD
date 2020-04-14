# Simplifying Goroutine Management

## about://errgroup.Group

> ```go
> type Group struct {
> // contains filtered or unexported fields
> }
> A Group is a collection of goroutines working on subtasks that are part of the same overall task.
>
> A zero Group is valid and does not cancel on error.

## Use Case

Removing redundant and error-prone boilerplate to handle the lifecycle of (possibly interconnected) goroutines, such as:

```go
    // omitting context and signal handling
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// do something
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// do something else
	}()

	wg.Wait()
```

> **Note:** An alternative package with a slightly different API (and use case) is [`oklog/run.Group`](https://blog.gopheracademy.com/advent-2017/run-group/).