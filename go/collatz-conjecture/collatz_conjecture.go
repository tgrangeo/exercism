package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	step := 0
	if n <= 0 {
		return 0, errors.New("negative input")
	}
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		step++
	}
	return step, nil
}
