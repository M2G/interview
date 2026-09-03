package main

import (
	"context"
	"database/sql"
	"fmt"
	"interview/db"
	"interview/infrastructure"
	"log"
)

/*
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
*/

func main() {
	sqlDB, err := sql.Open("postgres", "postgres://user:password@localhost/madb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	queries := db.NewQueries(sqlDB)
	repo := infrastructure.NewSQLCRepo(queries)
	result, err := repo.GetJobById(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
