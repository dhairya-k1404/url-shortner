package cache

import (
	"fmt"
	"url-shortner/database"
	"url-shortner/models"
)

func SetMapping(shortCode, longURL string) error {
	mapping := models.URLMapping{
		ShortCode: shortCode,
		LongURL:   longURL,
	}
	if err := database.DB.Create(&mapping).Error; err != nil {
		return fmt.Errorf("failed to store mapping in PostgreSQL: %w", err)
	}
	if err := RedisClient.Set(Ctx, shortCode, longURL, cacheTTL).Err(); err != nil {
		return fmt.Errorf("failed to cache mapping in Redis: %w", err)
	}
	return nil
}

// GetMapping retrieves a URL mapping, checking Redis first, then PostgreSQL.
func GetMapping(shortCode string) (string, bool) {
	longURL, err := RedisClient.Get(Ctx, shortCode).Result()
	if err == nil {
		return longURL, true
	}
	var mapping models.URLMapping
	if err := database.DB.First(&mapping, "short_code = ?", shortCode).Error; err != nil {
		return "", false
	}
	_ = RedisClient.Set(Ctx, shortCode, mapping.LongURL, cacheTTL).Err()
	return mapping.LongURL, true
}
