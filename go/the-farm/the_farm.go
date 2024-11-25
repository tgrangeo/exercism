package thefarm

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, n int)(float64,error){
	amount, err := fodderCalculator.FodderAmount(n)
	if (err != nil){
		return 0,err
	}
	factor, err :=  fodderCalculator.FatteningFactor();
	if (err != nil){
		return 0,err
	}
	return amount / factor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, n int){
	panic("")
}


// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(n int){
	panic("")
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
