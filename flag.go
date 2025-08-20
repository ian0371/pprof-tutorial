package main

import (
	"fmt"
	"net/http"
)

func flagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	flag := getFlag()

	fmt.Fprint(w, flag)
}

func getFlag() string {
	var result []byte

	for i := range funcArr {
		if funcArr[i]() {
			result = append(result, '1')
		} else {
			result = append(result, '0')
		}
	}

	var chars []byte
	for i := 0; i < len(result); i += 8 {
		var ch byte
		for j := 0; j < 8; j++ {
			if result[i+j] == '1' {
				ch = ch<<1 | 1
			} else {
				ch = ch << 1
			}
		}
		chars = append(chars, ch)
	}
	result = chars
	return string(result)
}

// do not modify this function!!
func encode(n int) bool {
	ret := true
	for i := 0; i < n; i++ {
		ret = !ret
	}
	return ret
}

// do not modify this function!!
func encode2(n int) bool {
	ret := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			ret = false
		}
	}
	return ret
}
