package main

import "fmt"

func main()  {
	var a int
	fmt.Scanf("%d", &a)

	for i := 0; i < a ; i ++ {
		var h, w, n, roomCount, floor int
		fmt.Scanf("%d %d %d", &h, &w, &n)

		if n % h == 0 {
			roomCount = n / h
		} else {
			roomCount = n / h + 1
		}

		floor = n % h

		if floor == 0 {
			floor = h
		}

		fmt.Println(100*floor + roomCount)
	}
}
