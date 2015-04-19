package main

import (
	"fmt"
	"math"
)

func main() {
	prime_limit := 2000000

	var primes [1000000]int
	primes[0] = 2
	count := 1
	num := 3
	inc := 2
	var sqrt_num int
	var isPrime bool
	for num < prime_limit {
		isPrime = true
		sqrt_num = int(math.Sqrt(float64(num)))
		for i := 0; i < count; i++ {
			if primes[i] > sqrt_num {
				break
			}
			if num % primes[i] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes[count] = num
			count++;
		}
		num += inc
	}
	sum := 0
	for i := 0; i < count; i++ {
		sum += primes[i]
	}
	fmt.Println(sum)
}
