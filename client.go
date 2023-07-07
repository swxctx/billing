package billing

import (
	"time"
)

var (
	client *Client
)

type Client struct {
	// client id
	clientId string
	// client secret
	clientSecret string
	// 包名
	packageName string
	// 刷新token
	refreshToken string
	// Debug 是否调试模式
	debug bool
	// TimeoutSecond 超时时间
	timeoutSecond int64
	// token
	accessToken string
	// token过期时间
	accessTokenExpire time.Time
}

// GetAccessToken
func GetAccessToken() (*AccessToken, error) {
	return client.getAccessToken()
}

// Verify
func Verify(productId, purchaseToken string) (*OrderInfo, error) {
	return client.verify(productId, purchaseToken)
}
