package bloom

import (
	"math/rand"

	"github.com/spaolacci/murmur3"
)

type (
	BloomFilter struct {
		m     uint64 // maximum size
		k     uint32 // number of hash functions
		b     *[]uint8
		seeds *[]uint32
	}
)

func New(m uint64, k uint32) *BloomFilter {
	b := make([]uint8, m)
	seeds := createRandomSeeds(k)
	return &BloomFilter{
		m:     m,
		k:     k,
		b:     &b,
		seeds: seeds,
	}
}

func (bf *BloomFilter) Add(data string) error {
	bf.setBits(data)
	return nil
}

func (bf *BloomFilter) setBits(data string) {
	for i := uint32(0); i < bf.k; i++ {
		h1, _ := calculateHash([]byte(data), (*bf.seeds)[i])
		index := h1 % bf.m
		(*bf.b)[index] = 1
	}
}

func calculateHash(encodedData []byte, seed uint32) (uint64, uint64) {
	return murmur3.Sum128WithSeed(encodedData, seed)
}

// createRandomSeeds creates an array of random seeds that will be used
// every time to calculate the hash number each time
func createRandomSeeds(k uint32) *[]uint32 {
	seedRange := 100000000
	seeds := make([]uint32, k)
	for i := uint32(0); i < k; i++ {
		seeds[i] = uint32(rand.Intn(seedRange))
	}
	return &seeds
}
