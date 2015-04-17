package main

import "fmt"

func main() {
	num := 600851475143
	largest := 0
	for num % 2 == 0 {
		num /= 2
	}
	
	p := 3
	for num != 1{
		for num % p == 0{
			largest = p
			num /= p
		}
		p += 2
	}
	fmt.Println(largest)
	
}
