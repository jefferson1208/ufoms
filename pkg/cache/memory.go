package cache

import (
	"errors"
	"strings"
	"time"
)

type MemoryClient struct {
	prefix             string
	cacheProvider      CacheProvider
	subscribedChannels []string
}

func NewMemoryProvider(config *CacheConfiguration) (ICache, error) {
	return &MemoryClient{prefix: config.CachePrefix, cacheProvider: MEMORY_CACHE, subscribedChannels: strings.Split(config.Channels, ";")}, nil
}

func (r *MemoryClient) Subscribe(outChan chan<- (*MessageChannel), readyChan chan<- struct{}) error {

	if len(r.subscribedChannels) == 0 {
		return errors.New("no channel provided for registration")
	}

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
