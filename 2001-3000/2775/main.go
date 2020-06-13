package main

import "fmt"

func main()  {
	var t, k, n int
	fmt.Scanf("%d", &t)

	for i := 0; i < t ; i ++ {
		fmt.Scanf("%d", &k)
		fmt.Scanf("%d", &n)

		fmt.Println(getTotalCount(k, n))
	}
}

func getTotalCount(k, n int) int {
	if k == 0 {
		return n
	}

	if n == 1 {
		return 1
	}

	return getTotalCount(k, n-1) + getTotalCount(k-1, n)
}
