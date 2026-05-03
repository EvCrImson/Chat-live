package Middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimitByRoute(rdb *redis.Client, limit int64, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()
		route := c.FullPath()

		key := "rate:" + ip + ":" + route

		count, err := rdb.Incr(c.Request.Context(), key).Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "error no cache"})
			c.Abort()
			return
		}

		if count == 1 {
			rdb.Expire(c.Request.Context(), key, window)
		}

		if count > limit {
			c.JSON(429, gin.H{
				"error": "Too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
