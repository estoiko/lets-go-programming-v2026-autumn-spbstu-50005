package main

import (
	"errors"
	"fmt"
)

func calculate(a, b int, op string) (float64, error) {
	switch op {
	case "+":
		return float64(a + b), nil
	case "-":
		return float64(a - b), nil
	case "*":
		return float64(a * b), nil
	case "/":
		if b == 0 {
			return 0, errors.New("Division by zero")
		}
		return float64(a) / float64(b), nil
	default:
		return 0, errors.New("Invalid operation")
	}
}

func main() {
	var a, b int
	var op string

	n, err := fmt.Scan(&a, &b, &op)
	if err != nil {
		switch n {
		case 0:
			fmt.Println("Invalid first operand")
		case 1:
			fmt.Println("Invalid second operand")
		}
		return
	}

	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
