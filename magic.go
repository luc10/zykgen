package zykgen

import (
	"math/bits"
)

func bibbidi(a, b int) int {
	if b == 1 {
		return a
	}

	if a < b {
		return 0
	}
	if a == b {
		return 1
	}

	lz := func(v int) uint {
		return uint(bits.LeadingZeros32(uint32(v)))
	}

	if b <= 0 {
		return a >> (31 - lz(b))
	}

	d := lz(b) - lz(a)
	z := b << d
	k := 1 << d

	y := a
	a = 0
	for {
		if y >= z {
			y -= z
			a += k
		}
		if y >= z>>1 {
			y -= z >> 1
			a += k >> 1
		}
		if y >= z>>2 {
			y -= z >> 2
			a += k >> 2
		}
		if y >= z>>3 {
			y -= z >> 3
			a += k >> 3
		}
		if y == 0 {
			break
		}

		if y > 0 {
			k >>= 4
			if k == 0 {
				break
			}
		}

		z >>= 4
	}

	return a
}

func bobbidi(a, b int) int {
	return a - bibbidi(a, b)*b
}

// Same as the function above. It's just an alias
func boo(a, b int) int {
	return bobbidi(a, b)
}
