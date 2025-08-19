package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	workers []chan struct{}
)

func goroutineHandler(w http.ResponseWriter, r *http.Request) {
	count := 10

	start := time.Now()
	log.Printf("Creating %d goroutines", count)

	var wg sync.WaitGroup
	wg.Add(count * 2) // We're creating 2 goroutines per iteration

	for i := 0; i < count; i++ {
		done := make(chan struct{})

		workers = append(workers, done)

		go func(id int) {
			defer wg.Done()
			select {
			case <-done:
				return
			case <-time.After(24 * time.Hour):
				return
			}
		}(i)

		go func(id int) {
			defer wg.Done()
			for {
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				_ = rand.Intn(1000)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	log.Printf("All goroutines finished after %v", time.Since(start))

	// Return HTTP response
	elapsed := time.Since(start)
	response := fmt.Sprintf("Created %d goroutines in %v. Worker started to wait for completion.", count, elapsed)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}
