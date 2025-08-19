package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // enable pprof HTTP endpoints
	"time"
)

func main() {
	// Setup HTTP handlers
	http.HandleFunc("/cpu", cpuBottleneckHandler)
	http.HandleFunc("/heap", heapBottleneckHandler)
	http.HandleFunc("/goroutine", goroutineHandler)
	http.HandleFunc("/flag", flagHandler)

	// Root handler with usage information
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"message": "pprof Tutorial HTTP Server",
			"endpoints": {
				"/cpu?iterations=N": "CPU-intensive task (default: 100000 iterations)",
				"/heap?size_mb=N": "Memory-intensive task (default: 50MB)",
				"/goroutine": "Goroutine-intensive task",
				"/leak": "Memory leak simulation for profiling",
				"/debug/pprof/": "Go pprof profiling endpoints"
			},
			"profiling": {
				"heap": "/debug/pprof/heap",
				"profile": "/debug/pprof/profile",
				"goroutine": "/debug/pprof/goroutine",
				"trace": "/debug/pprof/trace"
			},
			"timestamp": "%s"
		}`, time.Now().Format(time.RFC3339))
	})

	port := ":8080"
	log.Printf("Starting HTTP server on http://localhost%s", port)
	log.Printf("Profiling available at http://localhost%s/debug/pprof/", port)
	log.Printf("Usage information at http://localhost%s/", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
