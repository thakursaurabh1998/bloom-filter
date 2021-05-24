package bloom

import "testing"

func TestBasic(t *testing.T) {
	maxSize := uint64(10000)
	hashFunctionNum := uint16(4)
	f := New(maxSize, hashFunctionNum)
	if f.m != maxSize {
		t.Error("Max size doesn't match the set value")
	}

	if f.k != hashFunctionNum {
		t.Error("Count of hash functions to run doesn't match the set value")
	}
}
