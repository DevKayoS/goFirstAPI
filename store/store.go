package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface {
	GetFullUrl(ctx context.Context, code string) (string, error)
	SaveShortenedUrl(ctx context.Context, _url string) (string, error)
}

func NewStore(rdb *redis.Client) Store {
	return store{rdb}
}

func (s store) SaveShortenedUrl(ctx context.Context, _url string) (string, error) {
	var code string
	for range 5 {
		code = genCode()
		if err := s.rdb.HGet(ctx, "encurtador", code).Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}
			return "", fmt.Errorf("failed to get code from encurtador hasmap: %w", err)
		}
	}

	if err := s.rdb.HSet(ctx, "encurtador", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code from encurtador hasmap: %w", err)
	}

	return code, nil
}

func (s store) GetFullUrl(ctx context.Context, code string) (string, error) {
	fullUrl, err := s.rdb.HGet(ctx, "encurtador", code).Result()

	if err != nil {
		return "", fmt.Errorf("failed to get code from encurtador hasmap: %w", err)
	}

	return fullUrl, nil
}
