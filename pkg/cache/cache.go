package cache

import (
	"encoding/json"
	"time"
)

type CacheProvider string

const (
	REDIS_CACHE   CacheProvider = "REDIS"
	MEMORY_CACHE  CacheProvider = "MEMORY"
	UNKNOWN_CACHE CacheProvider = "UNKNOWN"
)

type MessageChannel struct {
	Channel      string
	Pattern      string
	Payload      string
	PayLoadSlice []string
}

func (i MessageChannel) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

type ICache interface {
	Subscribe(outChan chan<- (*MessageChannel), readyChan chan<- struct{}) error
	Publish(channel string, message interface{}) error
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Expire(key string, expiration time.Duration) error
	Ping() (string, error)
	GetProvider() string
}
