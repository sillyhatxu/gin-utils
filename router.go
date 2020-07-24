package ginutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(opts ...Option) (*gin.Engine, error) {
	config := &Config{
		OpenHealth: true,
	}
	for _, opt := range opts {
		opt(config)
	}
	router := gin.New()

	if config.OpenHealth {
		router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{"/health"},
		}))
		router.Use(gin.Recovery())
		router.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "UP", "message": "OK"})
		})
	}
	return router, nil
}
