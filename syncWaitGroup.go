package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(name string) {
	defer wg.Done()
	for i := 1; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, " : ", i)
	}
}

func main() {
	start := time.Now()
	wg.Add(1)
	go test("Test")
	wg.Add(1)
	go test("Test2")
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
