package main

import (
	"github.com/wty92911/car-server-demo-go/router"
	"log"
	"runtime/debug"
)

func main() {
	go func() {
		if r := recover(); r != nil {
			log.Printf("Service crashed with panic: %v\n", r)
			log.Printf("Stack trace: %s", string(debug.Stack()))
		}
	}()

	// 启动服务
	r := router.SetupRouter()
	r.Run(":3000")
}
