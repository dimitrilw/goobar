package primefactors

func PrimeFactors(n int) (res []int) {
	// keep dividing by 2 until we arrive at n is odd
	for n%2 == 0 {
		res = append(res, 2)
		n = n / 2
	}

	// n is now odd & cannot be even
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			res = append(res, i)
			n = n / i
		}
	}

	// add final n if it is not a 1
	if n > 2 {
		res = append(res, n)
	}

	return
}
