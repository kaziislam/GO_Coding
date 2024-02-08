package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("cannot divide %v by zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("======================")
	fmt.Printf("Diving %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(20, 5)
	test(50, 10)
	test(3, 0)
}