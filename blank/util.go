package main

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}


func Lcm(a int, b int) int {
	return a / Gcd(a, b)
}

func Gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a-b*(a/b)
	}

	return a
}

func Gcd64(a int64, b int64) int64 {
	for b != 0 {
		a, b = b, a-b*(a/b)
	}

	return a
}

func Sqr(a float64) float64 {
	return a * a
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func Sign(a int) int {
	if a > 0 {
		return 1
	}
	
	if a < 0 {
		return -1
	}
	
	return 0
}
