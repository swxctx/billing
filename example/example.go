package main

import (
	"fmt"

	"github.com/swxctx/billing"
	"github.com/xiaoenai/xmodel/redis"
)

func main() {
	err := billing.Init(&billing.Config{
		ClientId:     "",
		ClientSecret: "",
		PackageName:  "",
		RefreshToken: "",
		Debug:        true,
		Redis: &redis.Config{
			DeployType: "single",
			ForSingle: redis.SingleConfig{
				Addr: "127.0.0.1:6379",
			},
		},
	})
	if err != nil {
		fmt.Println("init err")
		return
	}

	// 获取token
	accessToken, err := billing.GetAccessToken()
	if err != nil {
		fmt.Printf("err-> %v", err)
		return
	}
	fmt.Printf("accessToken-> %v\n", accessToken)

	// 查询订单
	orderInfo, err := billing.Verify("", "")
	if err != nil {
		fmt.Printf("err-> %v", err)
		return
	}
	fmt.Printf("orderInfo-> %v", orderInfo)
}