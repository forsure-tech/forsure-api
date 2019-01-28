package server

import (
	"github.com/forsure-tech/forsure-web-api/internal/routes"
	"github.com/gin-gonic/gin"
)

// New Server
func New() *gin.Engine {

	r := gin.Default()
	routes.Set(r)

	return r

}
