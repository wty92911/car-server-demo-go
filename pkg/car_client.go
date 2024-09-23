package pkg

import (
	car "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/car/v20220110"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"os"
)

func NewCarClient() *car.Client {
	// 从环境变量中获取 Secret ID 和 Secret Key
	secretId := os.Getenv("SECRET_ID")
	secretKey := os.Getenv("SECRET_KEY")

	// 如果你使用其他方式获取配置，可以替换成相应的代码
	// secretId := Config.Get(DefaultKeys.SECRET_ID)
	// secretKey := Config.Get(DefaultKeys.SECRET_KEY)

	// 创建腾讯云认证对象
	credential := common.NewCredential(
		secretId,
		secretKey,
	)

	// 创建客户端配置对象
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 30
	cpf.SignMethod = "TC3-HMAC-SHA256"

	// 创建 CarClient 对象
	// cloud api region, for example: ap-guangzhou, ap-beijing, ap-shanghai
	client, err := car.NewClient(credential, "ap-tokyo", cpf)
	if err != nil {
		panic(err)
	}
	return client
}
