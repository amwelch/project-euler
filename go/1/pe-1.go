package main

import "fmt"

func main() {
	factors := [2]int{3, 5}
	max := 1000
	nums := make(map[int]bool)
	var found bool
	var candidate int
	for i := 1; i <= 1000; i++ {
		found = false
		for _,num := range factors {
			candidate = num*i
			if candidate < max {
				nums[candidate] = true
				found = true
			} 
		}
		if !found {
			break
		}
	}

	sum := 0;
	for num,_ := range nums {
		sum += num
	}
	fmt.Println(sum)
}
