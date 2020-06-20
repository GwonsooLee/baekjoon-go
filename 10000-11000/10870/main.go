package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(fi(n))
}

func fi(n int) int {
	if n  <= 1 {
		return n
	}

	return fi(n-1) + fi(n-2)
}
