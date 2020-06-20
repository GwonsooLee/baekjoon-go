package main

import (
	"fmt"
)

var (
	count = 0
	path = []string{}
)

func main()  {
	var n int
	fmt.Scanf("%d", &n)

	Hanoi(n, 1,2,3)

	fmt.Println(len(path))
	for _, p := range path {
		fmt.Println(p)
	}
}

func Hanoi(n, start, helper, end int)  {
	if n == 1 {
		count ++
		path = append(path, fmt.Sprintf("%d %d", start, end))
		return
	}

	Hanoi(n-1, start, end, helper)
	Hanoi(1, start, helper, end)
	Hanoi(n-1, helper, start, end)
}
