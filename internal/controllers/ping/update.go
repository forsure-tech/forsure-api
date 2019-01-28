package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Update function from pingo
func Update(c *gin.Context) {
	c.String(http.StatusOK, "put")
}
