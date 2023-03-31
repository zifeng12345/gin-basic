package redis

import (
	"fmt"
	"time"

	"nwd/shared/log"

	"github.com/go-redis/redis"
)

type Irediser interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Delete(key ...string) error
	Keys(key string) ([]string, error)
	Hincry(keyName, filedName string, number int64) int64
}

type RedisService struct {
	rdb *redis.Client
}

var rdbs Irediser

//
func Init(host, passwd string, port, db int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, db),
		Password: passwd,
		DB:       db,
	})

	if rdb != nil {
		log.GetLog().Error("Redis", "Main redis conect success")
		redis := RedisService{
			rdb: rdb,
		}

		rdbs = &redis
	} else {
		log.GetLog().Error("Redis", "Main redis init failed")
	}
}

func (r *RedisService) Hincry(keyName, filedName string, number int64) int64 {
	if number <= 0 {
		number = 1
	}
	res, err := r.rdb.HIncrBy(keyName, filedName, number).Result()
	if err != nil {
		log.GetLog().Error("Hincry failed", "Hincry failed: %s", err)
	}

	return res
}

func GetRedis() Irediser {
	return rdbs
}

func (r *RedisService) Get(key string) (string, error) {
	return r.rdb.Get(key).Result()
}

func (r *RedisService) Set(key string, value interface{}) error {
	_, err := r.rdb.Set(key, value, time.Minute).Result()
	return err
}

func (r *RedisService) Delete(key ...string) error {
	_, err := r.rdb.Del(key...).Result()
	return err
}

func (r *RedisService) Keys(key string) ([]string, error) {
	return r.rdb.Keys(key).Result()
}
