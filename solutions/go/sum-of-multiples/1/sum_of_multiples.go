package sumofmultiples

// SumMultiples calculates the unique sum of all multiples of divisors strictly less than limit.
func SumMultiples(limit int, divisors ...int) int {
	uniqueMultiples := make(map[int]struct{})

	for _, div := range divisors {
		if div == 0 {
			continue
		}
		for item := div; item < limit; item += div {
			uniqueMultiples[item] = struct{}{}
		}
	}
	ans := 0
	for val := range uniqueMultiples {
		ans += val
	}
	return ans
}
