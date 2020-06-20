package main

import (
	"fmt"
)

func main()  {
	var n, s int
	fmt.Scanf("%d", &n)
	a := []int{}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s)
		a = append(a, s)
	}

	//sort.Ints(a)

	InsertionSort(&a)

	for _, b := range a {
		fmt.Println(b)
	}
}

func InsertionSort(arr *[]int)  {
	for i := 1 ; i < len(*arr); i++ {
		target := (*arr)[i]
		for j := i-1 ; j >= 0 ; j-- {
			if target < (*arr)[j] {
				(*arr)[j+1] = (*arr)[j]
			} else {
				(*arr)[j+1] = target
				break
			}

			if j == 0 {
				(*arr)[j] = target
			}
		}
	}
}

