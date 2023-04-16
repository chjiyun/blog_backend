package middleware

import (
	"blog_backend/config"
	"context"

	"github.com/gin-gonic/gin"
)

// context传递
func SetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置超时 Context
		// timeoutContext, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		// defer cancel()
		c.Set("DB", config.DB.WithContext(context.Background()))
		// Add a context to the log entry.
		c.Set("Logger", config.Logger.WithContext(context.Background()))
		c.Next()
	}
}
