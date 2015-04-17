package main

import "fmt"

func main() {
	num_last := 1
	num_cur := 2
	holder := 0
	sum := 0
	for num_cur < 4000000 {
        	if num_cur % 2 == 0 {
			sum += num_cur
		}
		holder = num_cur
		num_cur += num_last
		num_last = holder
	}
	fmt.Println(sum)
}
