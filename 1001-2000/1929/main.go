package main

import (
	"fmt"
	"math"
)

func main()  {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)

	notPrimes := make([]bool, n+1)

	max := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < max+1; i++ {
		if notPrimes[i] == false {
			for j := i+i; j <= n; j+=i {
				notPrimes[j] = true
			}
		}
	}

	if m == 1 {
		m = 2
	}

	for i := m ; i < n+1 ; i++ {
		if ! notPrimes[i] {
			fmt.Println(i)
		}
	}
}
