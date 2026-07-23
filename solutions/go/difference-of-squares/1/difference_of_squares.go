package differenceofsquares

func SquareOfSum(n int) int {
    sum := int((n * (n + 1)) / 2)
    return sum * sum
}

func SumOfSquares(n int) int {
	ans := 0
    for i := 1; i <= n; i++ {
        ans += (i * i)
    }
    return ans
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
