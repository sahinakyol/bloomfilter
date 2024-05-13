package hash

func Fnv1aHash(s string) uint64 {
	const (
		offset64 uint64 = 14695981039346656037
		prime64  uint64 = 1099511628211
	)

	var hash = offset64

	for i := 0; i < len(s); i++ {
		hash ^= uint64(s[i])
		hash *= prime64
	}
	return hash
}

func StringSumHash(s string) uint64 {
	var sum uint64
	for _, c := range s {
		sum += uint64(c)
	}
	return sum
}

func Murmur3Hash(s string, seed uint32) uint32 {
	const (
		c1 uint32 = 0xcc9e2d51
		c2 uint32 = 0x1b873593
	)

	length := len(s)
	h1 := seed
	roundedEnd := length & ^3

	for i := 0; i < roundedEnd; i += 4 {
		k1 := uint32(s[i]) | uint32(s[i+1])<<8 | uint32(s[i+2])<<16 | uint32(s[i+3])<<24

		k1 *= c1
		k1 = (k1 << 15) | (k1 >> 17)
		k1 *= c2

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19)
		h1 = h1*5 + 0xe6546b64
	}

	// Tail
	var k1 uint32 = 0
	tailIndex := roundedEnd
	switch length & 3 {
	case 3:
		k1 ^= uint32(s[tailIndex+2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(s[tailIndex+1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(s[tailIndex])
		k1 *= c1
		k1 = (k1 << 15) | (k1 >> 17)
		k1 *= c2
		h1 ^= k1
	}

	// Finalization
	h1 ^= uint32(length)
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}
