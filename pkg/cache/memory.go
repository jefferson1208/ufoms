package cache

import (
	"time"

	"github.com/jefferson1208/ufoms/pkg/config"
)

type MemoryClient struct {
	prefix        string
	cacheProvider CacheProvider
}

func NewMemoryProvider(config *config.Configuration) (ICache, error) {
	return &MemoryClient{prefix: config.CachePrefix, cacheProvider: MEMORY_CACHE}, nil
}

func (r *MemoryClient) Subscribe(outChan chan<- (*MessageChannel), readyChan chan<- struct{}, channel ...string) error {
	return nil
}

func (r *MemoryClient) Publish(channel string, message interface{}) error {
	return nil
}

func (r *MemoryClient) Get(key string) (string, error) {
	return "", nil
}

func (r *MemoryClient) Set(key string, value interface{}, expiration time.Duration) error {
	return nil
}

func (r *MemoryClient) Expire(key string, expiration time.Duration) error {
	return nil
}

func (r *MemoryClient) Ping() (string, error) {
	return "", nil
}

func (r *MemoryClient) GetProvider() string {
	return string(r.cacheProvider)
}
