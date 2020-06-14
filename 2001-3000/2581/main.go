package main

import "fmt"

func main()  {
	var m, n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)

	count := 0
	sum := 0
	min := 0
	for i := m; i <= n ; i++ {
		if isPrime(i){
			if count == 0 {
				min = i
			}

			count++
			sum += i
		}
	}

	if count == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(sum)
		fmt.Println(min)
	}

}

func isPrime(n int) bool {
	count := 0
	for i := 1; i < n ; i++ {
		if n % i == 0 {
			count++
		}
	}

	if count == 1 {
		return true
	}
	return false
}
