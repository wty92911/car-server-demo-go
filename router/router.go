package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/car-server-demo-go/controller"
	"github.com/wty92911/car-server-demo-go/middleware"
	"log"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.Use(gin.Logger())
	log.Println("setup router")
	// 定义路由
	router.POST("/StartProject", controller.StartProject)
	router.POST("/StopProject", controller.StopProject)
	router.POST("/Enqueue", controller.Enqueue)
	router.POST("/Dequeue", controller.Dequeue)

	return router
}
