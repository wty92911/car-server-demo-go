package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/car-server-demo-go/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 定义路由
	router.POST("/StartProject", controller.StartProject)
	router.POST("/StopProject", controller.StopProject)
	router.POST("/Enqueue", controller.Enqueue)
	router.POST("/Dequeue", controller.Dequeue)

	return router
}
