package bloomfilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	hashes := []func(int) int{
		func(i int) int { return i % 5 },
		func(i int) int { return i % 10 },
	}

	bf := NewBloomFilter[int](100, hashes)

	bf.Add(10)
	bf.Add(20)
	bf.Add(30)

	if !bf.Contains(10) {
		t.Errorf("BloomFilter.Contains(10) = false; want true")
	}

	if !bf.Contains(20) {
		t.Errorf("BloomFilter.Contains(20) = false; want true")
	}

	if !bf.Contains(30) {
		t.Errorf("BloomFilter.Contains(30) = false; want true")
	}
}
