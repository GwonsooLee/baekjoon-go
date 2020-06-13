package main

import "fmt"

func main()  {
	var n, x, y int
	fmt.Scanf("%d", &n)

	for i := 0 ; i < n ; i ++ {
		fmt.Scanf("%d %d", &x, &y)

		fmt.Println(getMinimumSteps(x, y))
	}
}

func getMinimumSteps(x, y int) int {
	distance := y - x
	half := distance / 2

	sum := 0
	i := 0
	for sum < half {
		i++
		sum += i
	}

	if sum > half {
		sum -= i
		i--
	}

	k := 2*i
	if distance - 2*sum != 0 {
		if distance - 2*sum <= (i+1) {
			k++
		} else {
			k += 2
		}
	}

	return k
}
