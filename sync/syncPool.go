package sync

import (
	"bytes"
	"fmt"
	"sync"
)

// @see https://pkg.go.dev/sync#Pool

var bufferPool = sync.Pool{
	New: func() any {
		// Alloué seulement si le pool est vide
		return new(bytes.Buffer)
	},
}

func process(data string) string {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset() // IMPORTANT : on repart d'un état propre
	defer bufferPool.Put(buf)

	buf.WriteString("processed: ")
	buf.WriteString(data)

	return buf.String()
}

func SyncPool() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(process(fmt.Sprintf("item-%d", n)))
		}(i)
	}
	wg.Wait()
}
