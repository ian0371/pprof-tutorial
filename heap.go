package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var data [][]byte

// heapIntensiveTask creates memory pressure as another bottleneck
func heapIntensiveTask(sizeMB int) {
	// Allocate large slices to create memory pressure
	data = make([][]byte, sizeMB)
	for i := 0; i < sizeMB; i++ {
		data[i] = make([]byte, 1024*1024) // 1MB each
		// Fill with some data to prevent optimization
		for j := range data[i] {
			data[i][j] = byte(i + j)
		}
	}
}

// heapBottleneckHandler executes memory-intensive work
func heapBottleneckHandler(w http.ResponseWriter, r *http.Request) {
	// Get size parameter or default to 50MB
	sizeStr := r.URL.Query().Get("size_mb")
	sizeMB := 50
	if sizeStr != "" {
		if parsed, err := strconv.Atoi(sizeStr); err == nil {
			sizeMB = parsed
		}
	}

	start := time.Now()
	log.Printf("Starting memory-intensive task with %dMB allocation", sizeMB)

	heapIntensiveTask(sizeMB)

	duration := time.Since(start)
	log.Printf("Memory-intensive task completed in %v", duration)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
		"task": "memory_intensive",
		"allocated_mb": %d,
		"duration_ms": %d,
		"timestamp": "%s",
		"len_data": %d,
	}`, sizeMB, duration.Milliseconds(), time.Now().Format(time.RFC3339), len(data))
}
