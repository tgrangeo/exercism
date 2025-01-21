package prime

import (
	"errors"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	step := 4
	limit := int(math.Sqrt(float64(num))) + 1
	for i := 5; i <= limit; i += step {
		if num%i == 0 {
			return false
		}
		step = 6 - step
	}

	return true
}

func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}
	count := 0
	current := 1
	for count < n {
		current++
		if isPrime(current) {
			count++
		}
	}

	return current, nil
}
