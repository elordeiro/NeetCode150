package main

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := range 32 {
		res = res<<1 | num>>i&1
	}
	return res
}
