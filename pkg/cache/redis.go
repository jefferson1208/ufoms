package cache

import (
	"context"
	"crypto/tls"
	"errors"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	redisClient        *redis.Client
	prefix             string
	ctx                context.Context
	cacheProvider      CacheProvider
	subscribedChannels []string
}

func NewRedisProvider(config *CacheConfiguration) (ICache, error) {

	var tlsConfig *tls.Config

	if config.CacheEnableTls {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: config.CacheInsecureSkipVerify,
		}
	}

	redis := redis.NewClient(&redis.Options{
		Addr:      config.CacheHost,
		DB:        config.CacheDataBase,
		Password:  config.CachePassword,
		Username:  config.CacheUserName,
		TLSConfig: tlsConfig,
	})

	return &RedisClient{
		redisClient:        redis,
		prefix:             config.CachePrefix,
		ctx:                context.Background(),
		cacheProvider:      REDIS_CACHE,
		subscribedChannels: strings.Split(config.Channels, ";"),
	}, nil
}

func (r *RedisClient) Subscribe(outChan chan<- (*MessageChannel), readyChan chan<- struct{}) error {

	if len(r.subscribedChannels) == 0 {
		return errors.New("no channel provided for registration")
	}

	pubSub := r.redisClient.Subscribe(r.ctx, r.subscribedChannels...)
	defer pubSub.Close()
	defer pubSub.Unsubscribe(r.ctx, r.subscribedChannels...)

	_, err := pubSub.Receive(r.ctx)

	if err != nil {
		//log.Error("")
		return err
	}

	ch := pubSub.Channel()
	close(readyChan)

	for {
		msg, more := <-ch
		if !more {
			r.redisClient.Close()
			return nil
		}
		outChan <- &MessageChannel{Channel: msg.Channel, Pattern: msg.Pattern, Payload: msg.Payload, PayLoadSlice: msg.PayloadSlice}
	}
}

func (r *RedisClient) Publish(channel string, message interface{}) error {

	err := r.redisClient.Publish(r.ctx, channel, message).Err()

	if err != nil {
		//log.Error("")
	}

	return err
}

func (r *RedisClient) Get(key string) (string, error) {

	result, err := r.redisClient.Get(r.ctx, r.prefix+key).Result()
	return result, err
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {

	_, err := r.redisClient.Set(r.ctx, r.prefix+key, value, expiration).Result()
	return err
}

func (r *RedisClient) Expire(key string, expiration time.Duration) error {

	_, err := r.redisClient.Expire(r.ctx, r.prefix+key, expiration).Result()
	return err
}

func (r *RedisClient) Ping() (string, error) {

	pong, err := r.redisClient.Ping(r.ctx).Result()
	return pong, err
}

func (r *RedisClient) GetProvider() string {
	return string(r.cacheProvider)
}
