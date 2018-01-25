package diffsquares

func SumOfSquares(n int) int {
	r := 0
	for i := 1; i <= n; i++ {
		r += i * i
	}
	return r
}

func SquareOfSums(n int) int {
	r := 0
	for i := 1; i <= n; i++ {
		r += i
	}
	return r * r
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
