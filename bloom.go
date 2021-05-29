package bloom

import (
	"math"
	"math/rand"

	"github.com/spaolacci/murmur3"
)

type (
	BloomFilter struct {
		m     uint64    // maximum size
		k     uint64    // number of hash functions
		b     *[]uint8  // bitarray
		seeds *[]uint32 // seeds array to create multiple hashes
	}
)

// New returns a BloomFilter instance
// n = number of instance
// p = expected probability of false positives
func New(n uint64, p float64) *BloomFilter {
	m, k := estimateFilterMetrics(float64(n), p)
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
	for i := uint64(0); i < bf.k; i++ {
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
func createRandomSeeds(k uint64) *[]uint32 {
	seedRange := 100000000
	seeds := make([]uint32, k)
	for i := uint64(0); i < k; i++ {
		seeds[i] = uint32(rand.Intn(seedRange))
	}
	return &seeds
}

// estimateFilterMetrics estimates the metrics
// to create the filter with the input values of
// number of elems and the expected false positive rate
// Reference taken from: https://hur.st/bloomfilter/
// n = max elements expected
// p = expected probability of false positives
// m = bit array size
// k = number of hash function
func estimateFilterMetrics(n, p float64) (uint64, uint64) {
	m := math.Ceil((n * math.Log(p)) / math.Log(1/math.Pow(2, math.Log(2))))
	k := math.Round((m / n) * math.Log(2))

	return uint64(m), uint64(k)
}
