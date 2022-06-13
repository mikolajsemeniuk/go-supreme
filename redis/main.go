package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

var ConversionToStringError = errors.New("Cannot convert to string error")

type RedisReader interface {
	RedisRead(string) error
}

type RedisWriter interface {
	RedisWrite() string
}

type RedisStorage struct {
	Client *redis.Client
}

func (s *RedisStorage) Read(key string, reader RedisReader) error {
	document, err := s.Client.Get(key).Result()
	if err != nil {
		return err
	}
	return reader.RedisRead(document)
}

func (s *RedisStorage) Write(key string, document RedisWriter, duration time.Duration) error {
	content := document.RedisWrite()
	if content == "" {
		return ConversionToStringError
	}
	return s.Client.Set(key, content, duration).Err()
}
