package main

import (
	"fmt"

	"github.com/swxctx/billing"
)

func main() {
	err := billing.Init(&billing.Config{
		ClientId:     "",
		ClientSecret: "",
		PackageName:  "",
		RefreshToken: "",
		Debug:        true,
		RedisAddr:    "127.0.0.1:6379",
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

	productId := ""
	pucharseToken := ""

	// 查询订单
	orderInfo, err := billing.Verify(productId, pucharseToken)
	if err != nil {
		fmt.Printf("err-> %v", err)
		return
	}
	fmt.Printf("orderInfo-> %v", orderInfo)

	if orderInfo.ConsumptionState == billing.CONSUME_STATE_UNCONFIRMED {
		// 确认订单
		if err := billing.Consume(productId, pucharseToken); err != nil {
			fmt.Printf("Consume, err-> %v", err)
		}
	}
}
