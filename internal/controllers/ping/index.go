package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index function from pingo
func Index(c *gin.Context) {
	c.String(http.StatusOK, "xoxo")
}
