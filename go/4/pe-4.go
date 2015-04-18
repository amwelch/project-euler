package main

import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
	start := 100
	end := 1000
	largest := 0
	candidate := 0
	var substr_a string
	var substr_b string
	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			candidate = i*j	
			candidate_str := strconv.Itoa(candidate)
			//fmt.Println(candidate_str)
			if len(candidate_str) % 2 == 0 {
				substr_a = candidate_str[:len(candidate_str) / 2]
				substr_b = candidate_str[len(candidate_str) / 2 :]
			} else {
				substr_a = candidate_str[:len(candidate_str) / 2]
				substr_b = candidate_str[len(candidate_str) / 2 + 1:]
			}
			//fmt.Println(substr_a, Reverse(substr_b))
			if substr_a == Reverse(substr_b) && candidate > largest {
				largest = candidate
			}
		}
	}
	fmt.Println(largest)
	
}
