package main

func getSum(a int, b int) int {
	var total, carry int
	for {
		total = a ^ b
		carry = (a & b) << 1
		a, b = total, carry
		total |= carry
		if carry == 0 {
			break
		}
	}
	return total
}
