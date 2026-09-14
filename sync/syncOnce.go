package sync

import (
	"fmt"
	"sync"
)

type Connection struct {
	dsn string
}

var (
	instance *Connection
	once     sync.Once
)

func GetConnection(dsn string) *Connection {
	once.Do(func() {
		fmt.Println("initialisation de la connexion...") // exécuté une seule fois
		instance = &Connection{dsn: dsn}
	})
	return instance
}

func Once() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			conn := GetConnection("postgres://localhost/db")
			fmt.Printf("goroutine %d utilise %p\n", n, conn)
		}(i)
	}

	wg.Wait()
}
