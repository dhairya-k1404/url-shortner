package cache

import (
	"context"
	"log"
	"time"
	"url-shortner/conf"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

const cacheTTL = 10 * time.Minute

// InitRedis initializes the Redis client.
func InitRedis() {
	redisURL := conf.REDIS_URL
	if redisURL == "" {
		log.Fatal("❌ REDIS_URL is not set")
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("❌ Failed to parse Redis URL: %v", err)
	}

	RedisClient = redis.NewClient(opt)

	if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis")
}
