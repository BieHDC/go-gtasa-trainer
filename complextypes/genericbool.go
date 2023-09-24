package complextypes

func FlipBool[T integer](b T) T {
	if b == T(0) {
		return T(1)
	}
	return T(0)
}

func FlipBoolInplace[T integer](b *T) {
	*b = FlipBool(*b)
}
