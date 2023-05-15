package cache

import "github.com/dgraph-io/ristretto"

func New() (*ristretto.Cache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per getCached buffer.
	})
	if err != nil {
		return nil, err
	}
	return cache, err
}
