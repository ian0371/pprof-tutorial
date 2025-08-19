package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

// cpuBottleneckHandler executes CPU-intensive work
func cpuBottleneckHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Printf("Could not create CPU profile: %v", err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Get iterations parameter or default to 100000
	iterationsStr := r.URL.Query().Get("iterations")
	iterations := 100000
	if iterationsStr != "" {
		if parsed, err := strconv.Atoi(iterationsStr); err == nil {
			iterations = parsed
		}
	}

	start := time.Now()
	log.Printf("Starting CPU-intensive task with %d iterations", iterations)

	result := cpuIntensiveTask(iterations)

	duration := time.Since(start)
	log.Printf("CPU-intensive task completed in %v", duration)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
		"task": "cpu_intensive",
		"iterations": %d,
		"result": %f,
		"duration_ms": %d,
		"timestamp": "%s"
	}`, iterations, result, duration.Milliseconds(), time.Now().Format(time.RFC3339))
}

// cpuIntensiveTask creates a mock bottleneck by performing CPU-intensive calculations
func cpuIntensiveTask(iterations int) float64 {
	var result float64
	for i := 0; i < iterations; i++ {
		// Perform some meaningless but CPU-intensive calculations
		x := rand.Float64() * 100
		result += math.Sin(x) * math.Cos(x) * math.Sqrt(x)

		// Add some more computation to make it heavier
		for j := 0; j < 1000; j++ {
			result += math.Log(float64(j+1)) / math.Pow(2, float64(j%10))
		}
	}
	return result
}
