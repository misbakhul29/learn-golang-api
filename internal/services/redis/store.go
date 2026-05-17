package redis

import (
	"context"
	"encoding/json"
	"time"

	redisv9 "github.com/redis/go-redis/v9"
)

type Store struct {
	rdb *redisv9.Client
}

func NewStore(rdb *redisv9.Client) *Store {
	return &Store{rdb: rdb}
}

func (s *Store) Set(key string, value string, expiration time.Duration) error {
	return s.rdb.Set(context.Background(), key, value, expiration).Err()
}

func (s *Store) Get(key string) (string, error) {
	return s.rdb.Get(context.Background(), key).Result()
}

func (s *Store) Del(key string) error {
	return s.rdb.Del(context.Background(), key).Err()
}

func (s *Store) SetObject(key string, obj any, expiration time.Duration) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return s.Set(key, string(data), expiration)
}

func (s *Store) GetObject(key string, target any) error {
	val, err := s.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), target)
}
