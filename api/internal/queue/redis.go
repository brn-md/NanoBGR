package queue

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(url string) (*RedisClient, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisClient{Client: client}, nil
}

func (r *RedisClient) PushTask(ctx context.Context, queueName string, payload string) error {
	return r.Client.LPush(ctx, queueName, payload).Err()
}

func (r *RedisClient) SetStatus(ctx context.Context, key string, status string) error {
	return r.Client.Set(ctx, key, status, 24*time.Hour).Err()
}

func (r *RedisClient) GetStatus(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
