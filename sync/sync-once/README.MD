# To `sync.Once`...or twice

## about://sync.Once

> ```go
> type Once struct {
>   // contains filtered or unexported fields
> }
> ```
> 
> Once is an object that will perform exactly one action.

## Why would I need to re-run `sync.Once`? 

In certain situations it might be required to reset the function passed to `Do(f func())`. This could happen if a temporary (non-fatal) error has occurred, or to react to a change in the environment. 

Is that a good idea? Read [here](https://github.com/golang/go/issues/25955) what the Go engineers have to say.

> **Note:** Of course, Mat Ryer has us covered as usual: [matryer/resync](https://github.com/matryer/resync)

## Simple Benchmark of `sync.Once` vs. `sync.Mutex`

```bash
# run benchmark without compiler optimizations
▶ go test -bench=. -benchmem -gcflags="-N -l -m"
[...]
goos: darwin
goarch: amd64
pkg: embano1/go-meetup-lej-04-2020/sync/sync-once
Benchmark_ClientOnce-16         83198002                12.5 ns/op             0 B/op          0 allocs/op
Benchmark_ClientLock-16         10769286                 104 ns/op             0 B/op          0 allocs/op
PASS
ok      embano1/go-meetup-lej-04-2020/sync/sync-once    4.176s
```

> **Note:** Please DO NOT use this library, for educational purposes only.