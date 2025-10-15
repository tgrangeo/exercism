package darts

func Score(x, y float64) int {
	if x*x+y*y <= 1 {
		return 10
	}
	if x*x+y*y <= 25 {
		return 5
	}
	if x*x+y*y <= 100 {
		return 1
	} else {
		return 0
	}
}
