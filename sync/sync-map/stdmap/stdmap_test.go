package stdmap

import (
	"math/rand"
	"sync"
	"testing"
)

// prevent compiler optimizations
var result int

func genInts(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func BenchmarkStoreStandardMap(b *testing.B) {
	nums := genInts(b.N)
	sm := New()
	b.ResetTimer()
	for _, v := range nums {
		sm.Store(v, v)
	}
}

func BenchmarkStoreSyncMap(b *testing.B) {
	nums := genInts(b.N)
	var syncm sync.Map
	b.ResetTimer()
	for _, v := range nums {
		syncm.Store(v, v)
	}
}

func BenchmarkDeleteRegular(b *testing.B) {
	nums := genInts(b.N)
	sm := New()
	for _, v := range nums {
		sm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		sm.Delete(v)
	}
}

func BenchmarkDeleteSync(b *testing.B) {
	nums := genInts(b.N)
	var syncm sync.Map
	for _, v := range nums {
		syncm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		syncm.Delete(v)
	}
}

func BenchmarkLoadRegularFound(b *testing.B) {
	nums := genInts(b.N)
	sm := New()
	for _, v := range nums {
		sm.Store(v, v)
	}

	currentResult := 0
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		currentResult, _ = sm.Load(nums[i])
	}
	result = currentResult
}

func BenchmarkLoadSyncFound(b *testing.B) {
	nums := genInts(b.N)
	var syncm sync.Map
	for _, v := range nums {
		syncm.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, ok := syncm.Load(nums[i])
		if ok {
			currentResult = r.(int)
		}
	}
	result = currentResult
}
