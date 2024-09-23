package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/wty92911/car-server-demo-go/router"
)

func main() {
	// 定义命令行参数
	secretID := flag.String("SECRET_ID", "", "Your Secret ID")
	secretKey := flag.String("SECRET_KEY", "", "Your Secret Key")

	// 解析命令行参数
	flag.Parse()

	// 检查参数是否为空
	if *secretID == "" || *secretKey == "" {
		log.Fatal("SECRET_ID and SECRET_KEY must be provided")
	}

	// 打印参数（可选）
	fmt.Printf("SECRET_ID: %s\n", *secretID)
	fmt.Printf("SECRET_KEY: %s\n", *secretKey)

	// 设置密钥
	os.Setenv("SECRET_ID", *secretID)
	os.Setenv("SECRET_KEY", *secretKey)
	// 启动服务
	r := router.SetupRouter()
	r.Run(":3000")
}
