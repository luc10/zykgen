package zykgen

import (
	"crypto/md5"
	"fmt"
)

// Transform an input char
func transform(c int32, a int32, b int32) byte {
	if c+a <= 25 && c+a >= 0 {
		return byte(c + b)
	}

	return byte(c)
}

func Wpa(serial string, length int, cocktail Cocktail) (key string) {
	var (
		buf []byte
		c   int
	)

	// The first step requires to transform each device serial
	// char under some conditions
	for _, c := range serial {
		buf = append(buf, transform(c, -97, -32))
	}

	// Compute two md5 digests and use first and second byte
	// to obtain an int value which will be used later as base index
	// value to access to charset bytes
	dgst := md5.Sum([]byte(fmt.Sprintf("%02xPSK_ra0", md5.Sum(buf))))
	base := int(dgst[0])<<8 | int(dgst[1])

	// Since the md5 digest is 16 bytes long we can't request
	// for a password longer than 16 chars
	if length > 16 {
		length = 16
	}
	if cocktail < Mojito || cocktail > Cosmopolitan {
		cocktail = Mojito
	}

	for i, q := 0, 1; i < length; i, q = i+1, q*2 {
		if bibbidi(bobbidi(base, q*2), q) == 1 {
			c = boo(int(dgst[i]), 26) + 65
		} else {
			c = boo(int(dgst[i]), 10) + 48
		}

		// Concat each byte
		key += string(cocktails[cocktail](int32(c), base))
	}

	return
}
