package collatzconjecture

func CollatzConjecture(n int) (int, error) {
	step := 0
	for n != 1{
		if (n  % 2 == 0){
			n = n / 2
		} else {
			n = n * 3 + 1
		}
		step++
	}
	return step, nil
}
