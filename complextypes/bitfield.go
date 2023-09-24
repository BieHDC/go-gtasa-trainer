package complextypes

/*
// Counting up the bits
const (
	A uint8 = 1 << (1 * iota)
	B
	C
	D
	E
	F
	G
	H
)

// if you need it reversed for some reason
A uint8 = 128 >> (1 * iota)
*/

// These are for use with already pre-shifted values
// not for position indicators
func SetBit[T integer](source T, flag T) T {
	return source | flag
}

func GetBit[T integer](source T, flag T) bool {
	return (source & flag) != 0
}

func ClearBit[T integer](source T, flag T) T {
	return source &^ flag
}

func FlipBit[T integer](source T, flag T) T {
	return source ^ flag
}
