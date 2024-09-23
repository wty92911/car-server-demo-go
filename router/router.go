package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/car-server-demo-go/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 定义路由
	router.POST("/StartProject", controllers.StartProject)
	router.POST("/StopProject", controllers.StopProject)
	router.POST("/Enqueue", controllers.Enqueue)
	router.POST("/Dequeue", controllers.Dequeue)

	return router
}
