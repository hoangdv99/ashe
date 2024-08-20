// hint: find nth fibonaci number

func climbStairs(n int) int {
	a := 0
	b := 1
	var c int

	if n == 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}