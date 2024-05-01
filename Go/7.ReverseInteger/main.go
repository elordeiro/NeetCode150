package main

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}
	res := 0
	for x > 9 {
		res = res*10 + x%10
		x /= 10
	}

	if res > 214748364 || (res == 214748364 && (isNegative && x > 8 || x > 7)) {
		return 0
	}
	res = res*10 + x
	if isNegative {
		return -res
	}
	return res
}
