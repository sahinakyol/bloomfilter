package main

import (
	"bloomfilter/hash"
	"fmt"
	"math"
)

type BloomFilter struct {
	bitset []bool // Slice representing the bit array
	size   uint64 // Size of the bit array
	k      uint64 // Number of hash functions
}

func NewBloomFilter(n uint64, p float64) *BloomFilter {
	m := optimalSize(n, p)   // Calculate the optimal size of the bit array
	k := optimalHashes(m, n) // Calculate the optimal number of hash functions
	return &BloomFilter{
		bitset: make([]bool, m), // Initialize the bit array with the calculated size
		size:   m,               // Set the size of the bit array
		k:      k,               // Set the number of hash functions
	}
}

// optimalSize calculates the optimal size of the bit array
func optimalSize(n uint64, p float64) uint64 {
	m := -float64(n) * math.Log(p) / (math.Ln2 * math.Ln2) // Formula to calculate the optimal size
	return uint64(math.Ceil(m))                            // Return the ceiling of the calculated size
}

// optimalHashes calculates the optimal number of hash functions
func optimalHashes(m, n uint64) uint64 {
	k := (float64(m) / float64(n)) * math.Ln2 // Formula to calculate the optimal number of hash functions
	return uint64(math.Ceil(k))               // Return the ceiling of the calculated number of hash functions
}

func (bf *BloomFilter) Add(item string) {
	for i := uint64(0); i < bf.k; i++ { // Apply each hashedData function
		hashedData := hash.Fnv1aHash(fmt.Sprintf("%d%s", i, item)) // Compute the i-th hashedData of the item
		index := hashedData % bf.size                              // Compute the index by taking modulo with the size of the bit array
		bf.bitset[index] = true                                    // Set the bit at the computed index to true
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	for i := uint64(0); i < bf.k; i++ { // Apply each hashedData function
		hashedData := hash.Fnv1aHash(fmt.Sprintf("%d%s", i, item)) // Compute the i-th hashedData of the item
		index := hashedData % bf.size                              // Compute the index by taking modulo with the size of the bit array
		if !bf.bitset[index] {                                     // If any bit is not set, the item is definitely not in the filter
			return false
		}
	}
	return true // If all bits are set, the item is possibly in the filter
}
