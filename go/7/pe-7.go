package main

import (
	"fmt"
	"math"
)

func main() {

	prime_num := 10001

	var primes [100000]int
	primes[0] = 2
	count := 1
	num := 3
	inc := 2
	var sqrt_num int
	var isPrime bool
	for count < prime_num {
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
	fmt.Println(primes[prime_num-1])
}
