package handler

import (
	"net/http"
	"time"

	"github.com/dgraph-io/ristretto"
	"go.uber.org/zap"
)

type idempotent struct {
	cache *ristretto.Cache
	ttl   time.Duration
	log   *zap.Logger
}

func newIdempotent(cache *ristretto.Cache, log *zap.Logger) *idempotent {
	const (
		ttl = time.Hour
	)
	return &idempotent{
		cache: cache,
		ttl:   ttl,
		log:   log.Named("idempotent"),
	}
}

const (
	headerIdempotencyKey = "Idempotency-Key"
)

type idempotentCached struct {
	Status   int
	Header   http.Header
	Response []byte
}

func (i *idempotent) getCached(key string) (*idempotentCached, bool) {
	if key != "" {
		if val, ok := i.cache.Get(key); ok {
			cached, ok := val.(idempotentCached)
			if !ok {
				i.log.DPanic("wrong type idempotentCached")
			} else {
				i.log.Debug("return idempotentCached")
				return &cached, true
			}
		}
	}
	return nil, false
}

func (i *idempotent) setCached(status int, header http.Header, body []byte) {
	key := header.Get(headerIdempotencyKey)
	if len(key) > 0 {
		cached := idempotentCached{
			Status:   status,
			Header:   header,
			Response: body,
		}
		i.cache.SetWithTTL(key, cached, 0, i.ttl)
	}
}
