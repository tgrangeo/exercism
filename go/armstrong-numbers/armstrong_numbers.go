package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	str := strconv.Itoa(n)
	res := 0
	for _, c := range str {
		digit, _ := strconv.Atoi(string(c))
		res += int(math.Pow(float64(digit), float64(len(str))))

	}
	return res == n
}
