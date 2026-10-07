package main

import "fmt"

const (
	minTemp        = 15
	maxTemp        = 30
	errorCode      = -1
	lessOrEqual    = "<="
	greaterOrEqual = ">="
)

type bounds struct {
	left, right int
}

func (b *bounds) change(sign string, temp int) {
	switch sign {
	case lessOrEqual:
		b.right = min(b.right, temp)
	case greaterOrEqual:
		b.left = max(b.left, temp)
	}
}

func (b *bounds) check() bool {
	return b.left <= b.right
}

func main() {
	var departments, workers int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Error: invalid departments count:", err)

		return
	}

	for range departments {
		temperatureBounds := bounds{minTemp, maxTemp}

		if _, err := fmt.Scan(&workers); err != nil {
			fmt.Println("Error: invalid workers count:", err)

			return
		}

		for range workers {
			var (
				sign string
				temp int
			)

			if _, err := fmt.Scan(&sign, &temp); err != nil {
				fmt.Println("Error: invalid state:", err)

				return
			}

			temperatureBounds.change(sign, temp)

			if temperatureBounds.check() {
				fmt.Println(temperatureBounds.left)
			} else {
				fmt.Println(errorCode)
			}
		}
	}
}
