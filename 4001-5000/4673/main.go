package main

import (
	"fmt"
	"math"
)

const MAX_NUM = 10000
func main()  {
	n := make([]int, MAX_NUM+50)
	for i := 1; i < MAX_NUM; i++ {
		n[d(i)-1] = 1
	}

	for i, v := range n {
		index := i + 1
		if v != 1 && index <= MAX_NUM {
			fmt.Println(index)
		}
	}
}

func d(n int) int {
	sum := 0
	indexCount := 0
	for n > 0 {
		a := n / 10
		left := n % 10
		sum += int(math.Pow10(indexCount)) * left
		sum += left

		indexCount += 1
		n = a
	}
	return sum
}
