package billing

const (
	// doc:  https://developers.google.com/android-publisher/api-ref/rest/v3/purchases.products?hl=zh-cn
	// refresh token
	refreshTokenApi = "https://accounts.google.com/o/oauth2/token"
	// verify order api
	verifyOrderApi = "https://androidpublisher.googleapis.com/androidpublisher/v3/applications/%s/purchases/products/%s/tokens/%s"
	// 确认消耗商品[消耗型商品]
	consumeApi = "https://androidpublisher.googleapis.com/androidpublisher/v3/applications/%s/purchases/products/%s/tokens/%s:consume"
	// 确定交易[非消耗型商品]
	acknowledgeApi = "https://androidpublisher.googleapis.com/androidpublisher/v3/applications/%s/purchases/products/%s/tokens/%s:acknowledge"
)

const (
	// 支付成功
	ORDER_STATUS_SUCCESS = 0
	// 支付取消
	ORDER_STATUS_CANCEL = 0
	// 支付待处理
	ORDER_STATUS_PENDING = 0
)

const (
	// 消耗状态未确认
	CONSUME_STATE_UNCONFIRMED = 0
	// 消耗状态已确认
	CONSUME_STATE_CONFIRMED = 1
)

const (
	// 非消耗商品未确认
	ACKNOWLERDGE_STATE_UNCONFIRMED = 0
	// 非小号商品已确认
	ACKNOWLERDGE_STATE_CONFIRMED = 1
)
