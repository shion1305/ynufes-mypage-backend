package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"ynufes-mypage-backend/svc/runner"
)

func main() {
	engine := gin.New()
	apiV1 := engine.Group("/api/v1")
	err := runner.Implement(apiV1, false)
	if err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
	err = engine.Run("localhost:1306")
	if err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
}
