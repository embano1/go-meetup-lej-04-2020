package rpc

import "testing"

var benchClientOnce *Client
var benchClientLock *Client

func Benchmark_ClientOnce(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// The loop body is executed b.N times total across all goroutines.
			// The number of goroutines defaults to GOMAXPROCS.
			c := NewClientOnce()
			benchClientOnce = c
		}
	})
}

func Benchmark_ClientLock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// The loop body is executed b.N times total across all goroutines.
			// The number of goroutines defaults to GOMAXPROCS.
			c := NewClientLock()
			benchClientLock = c
		}
	})
}
