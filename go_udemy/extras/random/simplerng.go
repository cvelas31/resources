package main

func simpleRngRand(seed uint32) uint32 {
	num := seed
	num ^= (num << 13)
	num ^= (num >> 17)
	num ^= (num << 5)
	return num
}
