package billing

import (
	"time"

	"github.com/xiaoenai/xmodel/redis"
)

var (
	redisClient *redis.Client
)

const (
	// defaultTokenCacheKey 默认的accesstoken缓存的key
	defaultTokenCacheKey = "billing:swxctx:cache:access_token"
)

var (
	// accessTokenKey accessToken 缓存key
	accessTokenKey string
)

// initCache
func initCache(redisAddr string, cacheKey ...string) error {
	// init redis client
	cli, err := redis.NewClient(&redis.Config{
		DeployType: "single",
		ForSingle: redis.SingleConfig{
			Addr: redisAddr,
		},
	})
	if err != nil {
		return err
	}
	redisClient = cli

	// init cache key
	accessTokenKey = defaultTokenCacheKey
	if len(cacheKey) > 0 && len(cacheKey[0]) > 0 {
		accessTokenKey = cacheKey[0]
	}
	return nil
}

// setAccessToken 设置access token
func setAccessToken(token string, duration time.Duration) error {
	return redisClient.Set(accessTokenKey, token, duration).Err()
}

// getAccessToken 获取access_token
func getAccessToken() (string, time.Duration, error) {
	res := redisClient.Get(accessTokenKey)
	if res.Err() == redis.Nil {
		return "", 0, nil
	}
	return res.Val(), redisClient.TTL(accessTokenKey).Val(), res.Err()
}
