package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64{
		return 0, errors.New("bad")
	}
	res := uint64(1)
	for i := 1; i < number;i++{
		res *= 2
	}
	return res, nil
}

func Total() uint64 {
	return 18446744073709551615 
}
