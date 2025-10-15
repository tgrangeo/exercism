package thefarm

import (
	"errors";
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, n int) (float64, error) {
	amount, err := fodderCalculator.FodderAmount(n)
	if err != nil {
		return 0, err
	}
	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount / float64(n) * factor, nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, n int)(float64,error){
		if n <= 0 {
			return 0, errors.New("invalid number of cows")
		}
	return DivideFood(fodderCalculator,n)
}

func ValidateNumberOfCows(n int) error{
	if n < 0 {
		return fmt.Errorf("%d cows are invalid: there are no negative cows", n)
	} else if n == 0 {
		return fmt.Errorf("0 cows are invalid: no cows don't need food")
	}
	return nil
}
