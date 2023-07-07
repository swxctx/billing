package billing

const (
	// doc:  https://developers.google.com/android-publisher/api-ref/rest/v3/purchases.products?hl=zh-cn
	// refresh token
	refreshTokenApi = "https://accounts.google.com/o/oauth2/token"
	// verify order api
	verifyOrderApi = "https://androidpublisher.googleapis.com/androidpublisher/v3/applications/%s/purchases/products/%s/tokens/%s"
)

const (
	// 支付成功
	ORDER_STATUS_SUCCESS = 0
	// 支付取消
	ORDER_STATUS_CANCEL = 0
	// 支付待处理
	ORDER_STATUS_PENDING = 0
)
