package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch{
        case balance < 0.0:
    		return 3.213
        case balance >= 0.0 && balance < 1000.0:
    		return 0.5
        case balance >= 1000.0 && balance < 5000.0:
    		return 1.621
        case balance >= 5000.0:
    		return 2.475
    }
	return 0.0
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := float64(InterestRate(balance))
    return balance * (rate/float64(100.0))
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    for i:= 0 ; ; i++{
        if (balance >= targetBalance){
            return i
        }
        balance = AnnualBalanceUpdate(balance)
    }
	return 0
}
