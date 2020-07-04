package tik

import "math/bits"

// ROTR(a, k) is a circular shift to the right of bit string a by k slots.
// ROTR is a right shift, where overflowing bits to the right are added back to the left, instead of zeros.
// ROTR(a, k) = ROTL(a, n−k)
func rotr(a uint64, k int) uint64 {
	return bits.RotateLeft64(a, -k)
}

// ROTL(a, k) is a circular shift to the left of bit string a by k slots.
// ROTL is a left shift, where overflowing bits to the left are added back to the right, instead of zeros.
func rotl(a uint64, k int) uint64 {
	return bits.RotateLeft64(a, k)
}

// ctz input cannot be zero.
func ctz(a uint64) int {
	return bits.TrailingZeros64(a)
}

func clz(a uint64) int {
	return bits.LeadingZeros64(a)
}

func fls(a uint64) int {
	return 64 - clz(a)
}
