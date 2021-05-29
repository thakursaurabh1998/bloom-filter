package bloom_v1

import (
	"math/rand"
	"testing"
)

func TestEstimateFilterMetrics(t *testing.T) {
	maxElems := float64(100000)
	falsePositiveRate := 0.0000001
	m, k := estimateFilterMetrics(maxElems, falsePositiveRate)

	// example and test metrics taken directly
	// from the reference website default values
	// https://hur.st/bloomfilter/?n=100000&p=1.0E-7
	expectedElemsSize := uint64(3354771)
	expectedHashFunctions := uint64(23)

	if m != expectedElemsSize {
		t.Error("Max size doesn't match the expected calculated value")
	}

	if k != uint64(expectedHashFunctions) {
		t.Error("Count of hash functions doesn't match the expected calculated value")
	}
}

func TestHashCalculation(t *testing.T) {
	word := "data"
	seed := uint32(rand.Intn(100000))
	h1, h2 := calculateHash([]byte(word), seed)

	h3, h4 := calculateHash([]byte(word), seed)

	if h1 != h3 || h2 != h4 {
		t.Error("calculateHash function doesn't produce same hashes on similar inputs")
	}
}

func TestBloomFilterAdd(t *testing.T) {
	expectedElemsSize := uint64(10000)
	falsePositiveRate := 0.0001
	f := New(expectedElemsSize, falsePositiveRate)

	err := f.Add("test")

	if err != nil {
		t.Error("Error in adding word to bloom filter")
	}
}
