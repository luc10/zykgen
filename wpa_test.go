package zykgen

import (
	"testing"
)

var data = map[string]string{
	"S090Y00000000": "UJ4NKUJ8KP",
}

func TestWpa(t *testing.T) {
	for s, p := range data {
		w := Wpa(s, 10, Cosmopolitan)
		if w != p {
			t.Fatalf("got %s instead of %s expected wpa key", w, p)
		}
	}
}
