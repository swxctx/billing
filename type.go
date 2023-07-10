package billing

import "encoding/json"

// AccessToken
type AccessToken struct {
	// token
	AccessToken string `json:"access_token"`
	// 有效时间[单位为秒，默认一小时]
	ExpiresIn int64 `json:"expires_in"`
	// 授权范围
	Scope string `json:"scope"`
	// token 类型[Bearer]
	TokenType string `json:"token_type"`
}

// refreshTokenRequest
type refreshTokenRequest struct {
	// refresh_token
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// marshalAccessToken
func marshalAccessToken(accessTokenInfo *AccessToken) (string, error) {
	tokenByte, err := json.Marshal(accessTokenInfo)
	if err != nil {
		return "", err
	}
	return string(tokenByte), nil
}

// unmarshalAccessToken
func unmarshalAccessToken(tokenStr string) (*AccessToken, error) {
	var (
		accessToken *AccessToken
	)
	if err := json.Unmarshal([]byte(tokenStr), &accessToken); err != nil {
		return nil, err
	}
	return accessToken, nil
}

// verifyRequest
type verifyRequest struct {
	AccessToken string `json:"access_token"`
}

// OrderInfo 订单查询信息
type OrderInfo struct {
	// 支付时间，毫秒
	PurchaseTimeMillis string `json:"purchaseTimeMillis"`
	// 支付状态[0: 已购买 1: 已取消 2: 待处理]
	PurchaseState int `json:"purchaseState"`
	// 应用内商品的消耗状态[0: 尚未消耗 1: 已使用]
	ConsumptionState int `json:"consumptionState"`
	// 开发者指定的字符串，其中包含关于订单的补充信息
	DeveloperPayload string `json:"developerPayload"`
	// 订单ID
	OrderID string `json:"orderId"`
	// 应用内商品的购买类型。
	// 仅当购买交易不是使用标准的应用内购买结算流程完成时，系统才会设置此字段。
	// 可能的值包括：0. 测试（即从许可测试帐号中购买的服务）1. 促销（即使用促销代码购买）2. 激励广告（即通过观看视频广告而不是付费）
	PurchaseType int `json:"purchaseType"`
	// 应用内商品的确认状态。可能的值包括：0. 尚未确认 1. 已确认
	AcknowledgementState int `json:"acknowledgementState"`
	// 表示 android publisher service 中的 inappPurchase 对象
	Kind string `json:"kind"`
	// 产品被授予时用户的 ISO 3166-1 alpha-2 结算区域代码
	RegionCode string `json:"regionCode"`
}

// consumeRequest
type consumeRequest struct {
	AccessToken string `json:"access_token"`
}

// acknowledgeRequest
type acknowledgeRequest struct {
	AccessToken string `json:"access_token"`
}
