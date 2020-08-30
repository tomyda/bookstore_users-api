package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping is a funct
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
