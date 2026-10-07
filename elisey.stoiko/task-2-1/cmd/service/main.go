package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var n, k int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		l, r := minTemp, maxTemp
		bad := false
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			var sign string
			var t int

			fmt.Scan(&sign, &t)

			if !bad {
				switch sign[0] {
				case '<':
					r = min(r, t)
				case '>':
					l = max(l, t)
				}

				if l > r {
					bad = true
				}
			}

			if bad {
				fmt.Println("-1")
			} else {
				fmt.Println(l)
			}
		}
	}
}
