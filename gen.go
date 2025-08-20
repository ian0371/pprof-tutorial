//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"text/template"
)

type Func struct {
	ID   int
	Expr string
}

func isEven(n int) bool {
	return n%2 == 0
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func listOfEvens() []int {
	ret := []int{}
	for i := 100; i < 10000; i++ {
		if isEven(i) {
			ret = append(ret, i)
		}
		if len(ret) > 300 {
			break
		}
	}
	return ret
}

func listOfPrimes() []int {
	ret := []int{}
	for i := 100; i < 10000; i++ {
		if isPrime(i) {
			ret = append(ret, i)
		}
		if len(ret) > 300 {
			break
		}
	}
	return ret
}

func main() {
	flagStr := "FLAG{d3bugg1ng_with_pprof}"
	flag := make([]bool, 0, len(flagStr)*8)
	for _, c := range flagStr {
		// Convert each character to its binary representation
		for i := 7; i >= 0; i-- {
			flag = append(flag, (c>>uint(i))&1 == 1)
		}
	}

	// Define your expressions (could be dynamic or loaded from elsewhere)
	nums := listOfEvens()
	pivot := 144
	nums = append(nums[:pivot], listOfPrimes()...)

	funcs := []Func{}
	for i := range flag {
		funcStr := ""
		num := nums[i]
		// make flag false
		if !flag[i] {
			num += 1
		}

		if i == 87 { // 87th bit is false
			num = 12345678901234567
		} else if i == 173 { // 173th bit is false
			num = 89999999999999999
		}

		op1 := rand.Intn(1e18)
		if op1 < 1e17 {
			op1 += 1e17
		}
		op2 := num - op1

		if i < pivot {
			funcStr = fmt.Sprintf("encode(%d+%d)", op1, op2)
		} else {
			funcStr = fmt.Sprintf("encode2(%d+%d)", op1, op2)
		}

		// inject anomalies
		funcs = append(funcs, Func{ID: i, Expr: funcStr})
	}

	tmpl := template.Must(template.ParseFiles("funcs.tmpl"))
	out, err := os.Create("funcs.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, funcs)
	if err != nil {
		log.Fatal(err)
	}
}
