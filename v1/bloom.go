package bloom_v1

import (
	"math"
	"math/rand"

	"github.com/spaolacci/murmur3"
)

type (
	bloomFilter struct {
		m     uint64    // maximum size
		k     uint64    // number of hash functions
		b     *[]uint8  // bitarray
		seeds *[]uint32 // seeds array to create multiple hashes
	}

	BloomFilter interface {
		Add(string) error
		Lookup(string) bool
	}
)

// New returns a BloomFilter instance
// n = number of instance
// p = expected probability of false positives
func New(n uint64, p float64) BloomFilter {
	m, k := estimateFilterMetrics(float64(n), p)
	b := make([]uint8, m)
	seeds := createRandomSeeds(k)
	return &bloomFilter{
		m:     m,
		k:     k,
		b:     &b,
		seeds: seeds,
	}
}

// Add adds the data to the bloom filter
func (bf *bloomFilter) Add(data string) error {
	bf.setBits(data)
	return nil
}

// Lookup returns a boolean value
// if returned true it signifies that data **MIGHT** be present
// but a false return signifies that data is **GUARANTEED** not present
func (bf *bloomFilter) Lookup(data string) bool {
	for _, seed := range *bf.seeds {
		index := bf.findIndexAccToSeed(data, seed)
		if (*bf.b)[index] == 0 {
			return false
		}
	}

	return true
}

func (bf *bloomFilter) findIndexAccToSeed(data string, seed uint32) uint64 {
	h1, _ := calculateHash([]byte(data), seed)
	return h1 % bf.m
}

func (bf *bloomFilter) setBits(data string) {
	for _, seed := range *bf.seeds {
		index := bf.findIndexAccToSeed(data, seed)
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
