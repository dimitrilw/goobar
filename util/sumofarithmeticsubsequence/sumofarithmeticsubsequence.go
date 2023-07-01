package sumofarithmeticsubsequence

// cSpell:ignore sumofarithmeticsubsequence

// given (1, 5) as args, function returns 1+2+3+4+5 = 15
// given (3, 5) as args, function returns 3+4+5 = 12
func SumOfArithmeticSubSequence(a, b int) int {
	if a > b {
		a, b = b, a
	}
	n := b - a + 1
	// (a+b)/2 = average value; n = number of values
	// func does *n before /2 to preserve int and avoid floats
	return (a + b) * n / 2
}
