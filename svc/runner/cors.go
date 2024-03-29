package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/pkg/setting"
)

type CORS struct {
	targetHost string
}

func NewCORS() CORS {
	config := setting.Get()
	host := fmt.Sprintf("%s%s%s",
		config.Application.Server.Frontend.Protocol,
		config.Application.Server.Frontend.Domain,
		config.Application.Server.Frontend.Port)
	return CORS{targetHost: host}
}

func (cr CORS) ConfigureCORS(rg *gin.RouterGroup) {
	rg.Use(cr.middleware())
	// this does absolutely nothing because OPTIONS request will be intercepted by the middleware,
	// but this is needed to listen for OPTIONS requests
	rg.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})
}

func (cr CORS) middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cr.targetHost)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}
