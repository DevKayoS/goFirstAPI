package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface{}

func NewStore(rdb *redis.Client) Store {
	return store{rdb}
}

func (s store) SaveShortenedUrl(ctx context.Context, _url string) (string, error) {
	code := genCode()

	if err := s.rdb.HGet(ctx, "encurtador", code).Err(); err != nil {
		return "", fmt.Errorf("failed to get code from encurtador hasmap: %w", err)
	}

	if err := s.rdb.HSet(ctx, "encurtador", code, _url); err != nil {
		return "", fmt.Errorf("failed to set code from encurtador hasmap: %w", err)
	}

	return code, nil
}
