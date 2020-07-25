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
	if config.SkipPaths == nil {
		config.SkipPaths = []string{"/health"}
	} else {
		config.SkipPaths = append(config.SkipPaths, "/health")
	}
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: config.SkipPaths,
	}))
	router.Use(gin.Recovery())
	if config.OpenHealth {
		relativePath := "/health"
		if config.ContextPath != "" {
			relativePath = "/" + config.ContextPath + relativePath
		}
		router.GET(relativePath, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "UP", "message": "OK"})
		})
	}
	return router, nil
}
