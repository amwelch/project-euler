package main

import "fmt"

func main() {
	start := 1
	end := 100
	sum := 0
	sum_squares := 0
	for i := start; i <= end; i++ {
		sum += i
		sum_squares += i*i
	}
	square_sums := sum*sum
	fmt.Println(square_sums - sum_squares)
}
