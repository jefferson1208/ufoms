package cache

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jefferson1208/ufoms/pkg/config"
)

type RedisClient struct {
	redisClient   *redis.Client
	prefix        string
	ctx           context.Context
	cacheProvider CacheProvider
}

func NewRedisProvider(config *config.Configuration) (ICache, error) {

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
		redisClient:   redis,
		prefix:        config.CachePrefix,
		ctx:           context.Background(),
		cacheProvider: REDIS_CACHE,
	}, nil
}

func (r *RedisClient) Subscribe(outChan chan<- (*MessageChannel), readyChan chan<- struct{}, channel ...string) error {

	pubSub := r.redisClient.Subscribe(r.ctx, channel...)
	defer pubSub.Close()
	defer pubSub.Unsubscribe(r.ctx, channel...)

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
