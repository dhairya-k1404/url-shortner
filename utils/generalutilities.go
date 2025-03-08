package utils

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const shortURLLength = 6

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateShortCode returns a random string of fixed length.
func GenerateShortCode() string {
	b := make([]byte, shortURLLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
