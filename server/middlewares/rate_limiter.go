package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens         int
	maxTokens      int
	mutex          sync.Mutex
	refillInterval time.Duration
	refillAmount   int
}

func NewRateLimiter(maxTokens int, refillInterval time.Duration, refillAmount int) *RateLimiter {
	rl := &RateLimiter{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillInterval: refillInterval,
		refillAmount:   refillAmount,
	}
	go rl.refillTokens()
	return rl
}

func (rl *RateLimiter) refillTokens() {
	ticker := time.NewTicker(rl.refillInterval)
	for {
		<-ticker.C
		rl.mutex.Lock()
		rl.tokens = min(rl.maxTokens, rl.tokens+rl.refillAmount)
		rl.mutex.Unlock()
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// RateLimiterMiddleware applies rate limiting to incoming requests
func RateLimiterMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rl.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
