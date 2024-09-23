package main

import (
	"github.com/wty92911/car-server-demo-go/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":3000") // 启动服务
}
