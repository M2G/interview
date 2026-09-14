package sync

import (
	"fmt"
	"sync"
	"time"
)

// @see https://pkg.go.dev/sync#WaitGroup

func fetch(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // décrémente le compteur à la sortie (même en cas de panic)

	// simulation d'un traitement
	time.Sleep(time.Duration(id*100) * time.Millisecond)
	results <- fmt.Sprintf("ressource %d récupérée", id)
}

func SyncWaitGroup() {
	var wg sync.WaitGroup
	results := make(chan string, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1) // incrément avant le lancement de la goroutine
		go fetch(i, &wg, results)
	}

	// goroutine séparée pour fermer le channel une fois tout terminé
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}
