package main

import "fmt"

func main() {
	var ok bool
	num := 21
	for {
		ok = true
		for i := 1; i <= 20; i++ {
			if num % i != 0{
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(num)
			break
		}
		num++
	}
}
