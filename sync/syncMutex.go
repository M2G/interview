package sync

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		counts: make(map[string]int),
	}
}

func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock() // libéré (même si panic)
	c.counts[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}

func SyncMutex() {
	counter := NewSafeCounter()
	var wg sync.WaitGroup

	// 100 goroutines incrémentent la même clé en parallèle
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment("requests")
		}()
	}

	wg.Wait()
	fmt.Println("Total:", counter.Value("requests"))
}
