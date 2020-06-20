/*
Speed issue exists..

*/
package main

import (
	"fmt"
	"time"
)

var board [3000][3000]int

func main()  {
	var n int
	fmt.Scanf("%d", &n)
	start := time.Now()
	fillBoard(n, 0, 0)
	end := time.Now()


	for i := 0; i < n ; i++ {
		for j := 0 ; j < n ; j++ {
			if board[i][j] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			//fmt.Print(board[i][j])
		}
		fmt.Println()
	}

	fmt.Println(end.Sub(start))
}

func fillBoard(n, x, y int) {
	if n == 1 {
		board[x][y] = 1
		return
	}

	v := n/3
	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			if !(i == 1 && j == 1) {
				fillBoard(v, x+v*i, y+v*j)
			}
		}
	}

	return
}
