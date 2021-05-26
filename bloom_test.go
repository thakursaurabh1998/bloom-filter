package bloom

import (
	"math/rand"
	"testing"
)

func TestBasic(t *testing.T) {
	maxSize := uint64(10000)
	hashFunctionNum := uint32(4)
	f := New(maxSize, hashFunctionNum)
	if f.m != maxSize {
		t.Error("Max size doesn't match the set value")
	}

	if f.k != hashFunctionNum {
		t.Error("Count of hash functions to run doesn't match the set value")
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
	maxSize := uint64(10000)
	hashFunctionNum := uint32(4)
	f := New(maxSize, hashFunctionNum)

	err := f.Add("test")

	if err != nil {
		t.Error("Error in adding word to bloom filter")
	}
}
