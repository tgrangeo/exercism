package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	res := 0
	if len(digits) < span{
		return 0, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	for i := range digits {
		if len(digits[i:]) >= span {
			sub := digits[i : i+span]
			temp := 1
			for _, x := range sub {
				if x < '0' || x > '9' {
					return 0, errors.New("digits input must only contain digits")
				}
				temp = temp * int(x - '0')
			}
			if res < temp {
				res = temp
			}
		}
	}
	return int64(res), nil
}
