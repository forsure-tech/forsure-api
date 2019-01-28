package routes

import (
	"github.com/forsure-tech/forsure-web-api/internal/controllers/ping"
	"github.com/gin-gonic/gin"
)

// Set Routes
func Set(r *gin.Engine) {

	// Ping test
	r.GET("/ping", ping.Index)
	r.PUT("/ping", ping.Update)

}
