package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("vim-go")

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("ping succeeded")

	fmt.Printf("idle connections: %d\n", pool.Stat().IdleConns())
	fmt.Printf("max connections: %d\n", pool.Stat().MaxConns())

	_, err = pool.Exec(context.Background(), "CREATE TABLE ORDERS (amount int, product_name varchar(10))")
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	products := []string{"apple", "banana", "orange", "pine tree"}
	for _, product := range products {
		wg.Add(1)
		go insertProduct(product, pool, &wg)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go countProducts(pool, &wg)
	}
	wg.Wait()

	fmt.Printf("idle connections: %d\n", pool.Stat().IdleConns())
	fmt.Printf("max connections: %d\n", pool.Stat().MaxConns())
}

func insertProduct(product string, pool *pgxpool.Pool, wg *sync.WaitGroup) {
	_, err := pool.Exec(context.Background(), "INSERT INTO ORDERS (amount, product_name) VALUES (1, $1)", product)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
	return
}

func countProducts(pool *pgxpool.Pool, wg *sync.WaitGroup) {
	var count int
	err := pool.QueryRow(context.Background(), "SELECT count(*) FROM ORDERS").Scan(&count)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("row count: %d\n", count)
	wg.Done()
}
