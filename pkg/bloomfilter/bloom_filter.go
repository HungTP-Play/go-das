package bloomfilter

type BloomFilter[T any] struct {
	size   int
	hashes []func(T) int
	bits   []uint64
}

// NewBloomFilter creates a new BloomFilter with the specified size and hashes.
//
// Parameters:
// - size: the size of the bloom filter
// - hashes: an array of hash functions
// Returns a pointer to the newly created BloomFilter.
func NewBloomFilter[T any](size int, hashes []func(T) int) *BloomFilter[T] {
	return &BloomFilter[T]{
		size:   size,
		hashes: hashes,
		bits:   make([]uint64, size),
	}
}

// Add adds an item to the BloomFilter.
//
// Parameters:
// - item: the item to be added
func (b *BloomFilter[T]) Add(item T) {
	for _, hash := range b.hashes {
		b.bits[hash(item)%b.size] = 1
	}
}

// Contains checks if the BloomFilter possibly contains the item.
//
// Parameters:
// - item: the item to be checked
// Returns a boolean indicating whether the item is possibly in the BloomFilter.
func (b *BloomFilter[T]) Contains(item T) bool {
	for _, hash := range b.hashes {
		if b.bits[hash(item)%b.size] == 0 {
			return false
		}
	}
	return true
}
