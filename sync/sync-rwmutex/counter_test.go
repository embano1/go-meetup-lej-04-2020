package counter

import (
	"sync"
	"testing"
)

func parallelize(t *testing.T, workers int, testFn func(t *testing.T)) {
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			testFn(t)
		}()
	}
	wg.Wait()
}

// panics with incorrect use of Unlock()
func TestCounter_Get(t *testing.T) {
	type fields struct {
		counter int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"counter set to 2", fields{counter: 2}, 2},
		{"counter set to 6", fields{counter: 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counter{
				counter: tt.fields.counter,
			}
			if got := c.Get(); got != tt.want {
				t.Errorf("Counter.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// works (if Get() panic is resolved) because of no concurrency
func TestCounter_Increment(t *testing.T) {
	type fields struct {
		iterations int
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"increment 100 times w/out concurrency", fields{iterations: 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			for i := 0; i < tt.fields.iterations; i++ {
				c.Increment()
			}
			if got := c.Get(); got != tt.want {
				t.Errorf("Counter.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// will spot the race condition
func TestCounter_IncrementConcurrent(t *testing.T) {
	type fields struct {
		concurrency int
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"increment by one with 100 concurrent workers", fields{concurrency: 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			testFn := func(t *testing.T) {
				c.Increment()
			}
			parallelize(t, tt.fields.concurrency, testFn)

			if got := c.Get(); got != tt.want {
				t.Errorf("Counter.Get() = %v, want %v", got, tt.want)
			}

		})
	}
}
