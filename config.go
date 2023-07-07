package billing

import (
	"github.com/xiaoenai/xmodel/redis"
)

// Config owechat配置中心
type Config struct {
	// client id
	ClientId string `yaml:"client_id"`
	// client secret
	ClientSecret string `yaml:"client_secret"`
	// 包名
	PackageName string `yaml:"package_name"`
	// 刷新token
	RefreshToken string `yaml:"refresh_token"`
	// redis缓存token的key(选填)
	CacheKey string `yaml:"cache_key"`
	// Debug 是否调试模式
	Debug bool `yaml:"debug"`
	// TimeoutSecond 超时时间
	TimeoutSecond int64 `yaml:"timeout_second"`
	// Redis 缓存的redis配置
	Redis *redis.Config `yaml:"redis"`
}

// Init
func Init(cfg *Config) error {
	// init client
	client = &Client{
		clientId:      cfg.ClientId,
		clientSecret:  cfg.ClientSecret,
		packageName:   cfg.PackageName,
		refreshToken:  cfg.RefreshToken,
		debug:         cfg.Debug,
		timeoutSecond: cfg.TimeoutSecond,
	}

	if client.timeoutSecond <= 0 {
		client.timeoutSecond = 20
	}

	// init cache redis
	if err := initCache(cfg.Redis, cfg.CacheKey); err != nil {
		return err
	}
	return nil
}
