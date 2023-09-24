package complextypes

// Somehow it doesnt like the constrains package right now
// fixme investiage in later go version
type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type integer interface {
	signed | unsigned
}
