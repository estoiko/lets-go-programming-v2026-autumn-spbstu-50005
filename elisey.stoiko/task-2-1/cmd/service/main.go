package main

import "fmt"

const (
	minTemp   = 15
	maxTemp   = 30
	errorCode = -1
)

type bounds struct {
	left, right int
}

func (b *bounds) change(sign string, t int) {
	switch sign {
	case "<=":
		b.right = min(b.right, t)
	case ">=":
		b.left = max(b.left, t)
	}
}

func (b bounds) check() bool {
	return b.left <= b.right
}

func main() {
	var n, k int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		temperatureBounds := bounds{minTemp, maxTemp}

		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			var sign string
			var t int

			fmt.Scan(&sign, &t)
			temperatureBounds.change(sign, t)

			if temperatureBounds.check() {
				fmt.Println(temperatureBounds.left)
			} else {
				fmt.Println(errorCode)
			}
		}
	}
}
