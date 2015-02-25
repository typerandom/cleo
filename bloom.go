package cleo

const (
	fnvBasis64 = uint64(14695981039346656037)
	fnvPrime64 = uint64((1 << 40) + 435)
	fnvMask64  = uint64(^uint64(0) >> 1)
	numBits    = 64
)

//The bloom filter of a word is 8 bytes in length
//and has each character added separately
func computeBloomFilter(s string) int {
	cnt := len(s)

	if cnt <= 0 {
		return 0
	}

	var filter int
	hash := uint64(0)

	for i := 0; i < cnt; i++ {
		c := s[i]

		//first hash function
		hash ^= uint64(0xFF & c)
		hash *= fnvPrime64

		//second hash function (reduces collisions for bloom)
		hash ^= uint64(0xFF & (c >> 16))
		hash *= fnvPrime64

		//position of the bit mod the number of bits (8 bytes = 64 bits)
		bitpos := hash % numBits

		if bitpos < 0 {
			bitpos += numBits
		}

		filter = filter | (1 << bitpos)
	}

	return filter
}
