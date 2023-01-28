package ping

import "github.com/gin-gonic/gin"

// Ping @Summary Ping the server
// @Description Ping the server
// @Tags Ping
// @Accept  json
// @Produce  string
// @Success 200 {object} _root.PingResponse
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(200, "%s", "pong")
}
