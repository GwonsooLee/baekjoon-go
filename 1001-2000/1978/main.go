package main

import "fmt"

func main()  {
	var n, a int
	fmt.Scanf("%d", &n)

	count := 0
	for i := 0; i < n ; i++ {
		fmt.Scanf("%d", &a)
		if isPrime(a){
			count++
		}
	}

	fmt.Println(count)
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
