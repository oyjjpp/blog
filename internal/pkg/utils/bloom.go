package utils

import (
	"github.com/bits-and-blooms/bloom/v3"
)

// NewBloomFilter creates a Bloom filter for n expected items and a target false positive rate.
func NewBloomFilter(n uint, f float64) *bloom.BloomFilter {
	m, k := bloom.EstimateParameters(n, f)
	return bloom.New(m, k)
}

// LoadUserIDsToBloom adds user IDs to the Bloom filter in batch.
func LoadUserIDsToBloom(bf *bloom.BloomFilter, userIDs []string) {
	for _, uid := range userIDs {
		bf.AddString(uid)
	}
}

// BloomToBytes serializes a Bloom filter for persistence or transport.
func BloomToBytes(bf *bloom.BloomFilter) ([]byte, error) {
	return bf.MarshalBinary()
}

// BytesToBloom deserializes a Bloom filter from persisted bytes.
func BytesToBloom(data []byte) (*bloom.BloomFilter, error) {
	bf := &bloom.BloomFilter{}
	return bf, bf.UnmarshalBinary(data)
}

// CheckUserInBloom returns true when the user may exist and false when it definitely does not.
func CheckUserInBloom(bf *bloom.BloomFilter, userID string) bool {
	return bf.TestString(userID)
}
