package filter

import (
	"hash"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type bloomFilter struct {
	bitset    []bool
	k         uint          // number of hash functions
	m         uint          // size of the bloom filter bitset
	n         uint          // number of elements in the bloom filter
	hashFuncs []hash.Hash64 // hash functions
}

func newFilter(size int) *bloomFilter {
	return &bloomFilter{
		bitset:    make([]bool, size),
		k:         3,
		m:         uint(size),
		n:         uint(0),
		hashFuncs: []hash.Hash64{murmur3.New64(), fnv.New64(), fnv.New64a()},
	}
}

func (bf *bloomFilter) add(item []byte) {
	hashes := bf.hashValues(item)
	for i := uint(0); i < bf.k; i++ {
		position := uint(hashes[i]) % bf.m
		bf.bitset[position] = true
	}
	bf.n++
}

func (bf *bloomFilter) test(item []byte) bool {
	exists := true
	hashes := bf.hashValues(item)
	for i := uint(0); i < bf.k; i++ {
		position := uint(hashes[i]) % bf.m
		if !bf.bitset[position] {
			exists = false
			break
		}
	}
	return exists
}

func (bf *bloomFilter) hashValues(item []byte) []uint64 {
	var result []uint64
	for _, hashFunc := range bf.hashFuncs {
		hashFunc.Write(item)
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}
	return result
}
