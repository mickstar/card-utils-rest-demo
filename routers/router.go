package routers

import (
	"card-utils-rest-server/routers/pan"
	"card-utils-rest-server/routers/ping"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", ping.Ping)
	router.POST("/pan/mask", pan.MaskPan)
	router.GET("/pan/generate", pan.GenerateRandomPan)
	router.POST("/pan/validate", pan.ValidatePan)
	return router
}
