package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main()  {
	//var n, s int
	var n int

	fmt.Scanf("%d", &n)
	a := [1000000]int{}
	//for i := 0; i < n; i++ {
	//	fmt.Scanf("%d", &s)
	//	a[i] = s
	//}

	for i := n; i > 0; i-- {
		a[n-i] = i
	}
	start := time.Now()
	ret := MergeSort(a[0:n], 0, n-1)

	strs := []string{}

	for i := 0 ; i < n ; i++ {
		strs = append(strs, strconv.Itoa(ret[i]))
	}

	test := strings.Join(strs, "\n")
	fmt.Println(test)

	fmt.Println(time.Now().Sub(start))

}

func MergeSort(arr []int, start, end int) []int {
	sub := end-start
	if sub == 0 {
		return []int{arr[start]}
	}
	if sub == 1 {
		if arr[start] > arr[end] {
			return []int{arr[end], arr[start]}
		} else {
			return []int{arr[start], arr[end]}
		}
	}

	//Left mergeSort
	middle := start + (sub / 2)
	left := MergeSort(arr, start, middle)
	right := MergeSort(arr, middle+1, end)

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	ret := []int{}

	for leftIndex < len(left) || rightIndex < len(right) {
		if leftIndex == len(left) && rightIndex < len(right) {
			ret = append(ret, right[rightIndex])
			rightIndex++
		} else if leftIndex < len(left) && rightIndex == len(right) {
			ret = append(ret, left[leftIndex])
			leftIndex++
		} else if left[leftIndex] < right[rightIndex] {
			ret = append(ret, left[leftIndex])
			leftIndex++
		} else {
			ret = append(ret, right[rightIndex])
			rightIndex++
		}
	}

	return ret
}
