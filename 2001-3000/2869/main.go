package main

import "fmt"

func main()  {
	var a, b, v, days, step int
	fmt.Scanf("%d %d %d", &a, &b, &v)
	step = a-b

	if v % step == 0 {
		days = v / step
	} else {
		days = (v / step) + 1
	}

	for {
		if (days-1) * step + b >= v {
			days --
		}else {
			break
		}
	}

	fmt.Println(days)
}


