package zykgen

type Cocktail int
type CocktailFunction func(needle int32, base int) byte

const (
	Mojito Cocktail = iota
	Negroni
	Cosmopolitan
)

var cocktails = map[Cocktail]CocktailFunction{
	Mojito:       mojito,
	Negroni:      negroni,
	Cosmopolitan: cosmopolitan,
}

var ingredients = map[Cocktail]struct {
	haystack []byte
	charset  []byte
}{
	Mojito: {
		haystack: []byte{
			'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'i', 'l', 'o', 's', 'a', 'b',
			'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 't', 'u', 'v',
			'w', 'x', 'y', 'z', '1', '2', '5', '6', '9', '0', 'I', 'O', 'S', 'Z', '3', '4',
			'7', '8', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P',
			'Q', 'R', 'T', 'U', 'V', 'W', 'X', 'Y', '1', '2', '5', '6', '9', '0', 'I', 'O',
			'S', 'V', 'W', 'Z', '3', '4', '7', '8', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
			'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'T', 'U', 'X', 'Y',
		},
		charset: []byte{
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 't',
			'u', 'v', 'w', 'x', 'y', 'z', '1', '2', '5', '6', '9', '0', 'I', 'O', 'S', 'Z',
			'3', '4', '7', '8', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
			'N', 'P', 'Q', 'R', 'T', 'U', 'V', 'W', 'X', 'Y', '1', '2', '5', '6', '9', '0',
			'I', 'O', 'S', 'V', 'W', 'Z', '3', '4', '7', '8', 'A', 'B', 'C', 'D', 'E', 'F',
			'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'T', 'U', 'X', 'Y',
		},
	},
	Negroni: {
		haystack: []byte{
			'1', '2', '5', '6', '9', '0', 'I', 'O', 'S', 'Z', '3', '4', '7', '8', 'A', 'B',
			'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'T', 'U',
			'V', 'W', 'X', 'Y', '1', '2', '5', '6', '9', '0', 'I', 'O', 'S', 'V', 'W', 'Z',
			'3', '4', '7', '8', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
			'N', 'P', 'Q', 'R', 'T', 'U', 'X', 'Y',
		},
		charset: []byte{
			'3', '4', '7', '8', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
			'N', 'P', 'Q', 'R', 'T', 'U', 'V', 'W', 'X', 'Y', '1', '2', '5', '6', '9', '0',
			'I', 'O', 'S', 'V', 'W', 'Z', '3', '4', '7', '8', 'A', 'B', 'C', 'D', 'E', 'F',
			'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'T', 'U', 'X', 'Y',
		},
	},
	Cosmopolitan: {
		haystack: []byte{
			'1', '2', '5', '6', '9', '0', 'I', 'O', 'S', 'V', 'W', 'Z', '3', '4', '7', '8',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'K', 'L', 'J', 'M', 'N', 'P', 'Q', 'R',
			'T', 'U', 'X', 'Y',
		},
		charset: []byte{
			'3', '4', '7', '8', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
			'N', 'P', 'Q', 'R', 'T', 'U', 'X', 'Y',
		},
	},
}

func pick(haystack, charset []byte, needle byte, base, max, v int) byte {
	for i := 0; i < max; i++ {
		if haystack[i] == byte(needle) {
			return charset[bobbidi(base+i, v)]
		}
	}

	return byte(needle)
}

// Mojito method requires needle to be transformed before
// being passed to pick function
func mojito(needle int32, base int) byte {
	return pick(
		ingredients[Mojito].haystack,
		ingredients[Mojito].charset,
		transform(needle, -65, 32),
		base,
		14, 22,
	)
}

func negroni(needle int32, base int) byte {
	return pick(
		ingredients[Negroni].haystack,
		ingredients[Negroni].charset,
		byte(needle),
		base,
		10, 26,
	)
}

func cosmopolitan(needle int32, base int) byte {
	return pick(
		ingredients[Cosmopolitan].haystack,
		ingredients[Cosmopolitan].charset,
		byte(needle),
		base,
		12, 24,
	)
}
