package bloom

type (
	BloomFilter struct {
		m uint64 // maximum size
		k uint16 // number of hash functions
		b *[]uint8
	}
)

func New(m uint64, k uint16) *BloomFilter {
	b := make([]uint8, m)
	return &BloomFilter{
		m: m,
		k: k,
		b: &b,
	}
}
