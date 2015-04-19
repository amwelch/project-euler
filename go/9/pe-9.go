package main

import (
	"fmt"
)

func main() {
	var j int;
	var k int;
	var i_sq int;
	var j_sq int;
	var k_sq int;
	for i := 1; i < 500; i++ {
		i_sq = i*i
		for j = i+1; j < 1000; j++ {
			if i + j > 1000 {
				continue
			}
			k = 1000 - (i + j)
			j_sq = j*j;
			k_sq = k*k;
			if i_sq + j_sq == k_sq {
				fmt.Println(i*j*k)
			}
		}
	}
}
