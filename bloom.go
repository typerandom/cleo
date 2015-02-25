package cleo

//Used for the bloom filter
const (
	FNV_BASIS_64 = uint64(14695981039346656037)
	FNV_PRIME_64 = uint64((1 << 40) + 435)
	FNV_MASK_64  = uint64(^uint64(0) >> 1)
	NUM_BITS     = 64

	FNV_BASIS_32 = uint32(0x811c9dc5)
	FNV_PRIME_32 = uint32((1 << 24) + 403)
	FNV_MASK_32  = uint32(^uint32(0) >> 1)
)

//The bloom filter of a word is 8 bytes in length
//and has each character added separately
func ComputeBloomFilter(s string) int {
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
		hash *= FNV_PRIME_64

		//second hash function (reduces collisions for bloom)
		hash ^= uint64(0xFF & (c >> 16))
		hash *= FNV_PRIME_64

		//position of the bit mod the number of bits (8 bytes = 64 bits)
		bitpos := hash % NUM_BITS
		if bitpos < 0 {
			bitpos += NUM_BITS
		}
		filter = filter | (1 << bitpos)
	}

	return filter
}
