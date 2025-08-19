package main

import (
	"fmt"
	"net/http"
)

func flagHandler(w http.ResponseWriter, r *http.Request) {
	inputs := []int{10, 20, 30, 40, 42, 50}
	for _, n := range inputs {
		fmt.Printf("Secret for %d: %d\n", n, encode(n))
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Done")
}

func encode(n int) int {
	result := 0
	for i := 0; i < 100_000; i++ {
		result += computeLayer1(n, i)
	}
	return result
}

func computeLayer1(n, i int) int {
	sum := 0
	for j := 0; j < 10; j++ {
		sum += computeLayer2(n, i, j)
	}
	return sum
}

func computeLayer2(n, i, j int) int {
	return (i * j * n) % 123
}
