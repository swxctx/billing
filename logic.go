package billing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/swxctx/ghttp"
)

func (c *Client) getAccessToken() (*AccessToken, error) {
	// 从缓存获取
	tokenStr, ttl, err := getAccessToken()
	if err != nil {
		return nil, err
	}
	if len(tokenStr) > 0 {
		accessToken, err := unmarshalAccessToken(tokenStr)
		if err != nil {
			return nil, err
		}
		c.accessToken = accessToken.AccessToken
		c.accessTokenExpire = time.Now().Add(ttl)
		return accessToken, nil
	}

	// 重新从远程获取
	requestParams := refreshTokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: c.refreshToken,
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
	}

	// new request
	req := ghttp.Request{
		Url:       refreshTokenApi,
		Body:      requestParams,
		Method:    "POST",
		ShowDebug: client.debug,
		Timeout:   time.Duration(c.timeoutSecond) * time.Second,
	}
	req.AddHeader("Content-Type", "application/json")

	// send request
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get access token http status code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// resp
	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var (
		tokenResp *AccessToken
	)

	err = json.Unmarshal(respBs, &tokenResp)
	if err != nil {
		return nil, err
	}

	if tokenResp == nil {
		return nil, fmt.Errorf("get access token is nil for remote")
	}

	// 重新存储到redis
	setAccessToken(string(respBs), time.Duration(tokenResp.ExpiresIn-100)*time.Second)

	c.accessToken = tokenResp.AccessToken
	c.accessTokenExpire = time.Now().Add(time.Duration(tokenResp.ExpiresIn-100) * time.Second)
	return tokenResp, nil
}

// verify 校验订单是否正常
func (c *Client) verify(productId, purchaseToken string) (*OrderInfo, error) {
	// 检查刷新 access token
	if len(c.accessToken) <= 0 || c.accessTokenExpire.Unix() < time.Now().Unix() {
		// 刷新token
		_, err := c.getAccessToken()
		if err != nil {
			return nil, err
		}
	}

	// 校验订单
	requestParams := verifyRequest{
		AccessToken: c.accessToken,
	}

	// new request
	req := ghttp.Request{
		Url:       fmt.Sprintf(verifyOrderApi, c.packageName, productId, purchaseToken),
		Query:     requestParams,
		Method:    "GET",
		ShowDebug: client.debug,
		Timeout:   time.Duration(c.timeoutSecond) * time.Second,
	}
	req.AddHeader("Content-Type", "application/json")

	// send request
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("verify order http status code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// resp
	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var (
		orderInfo *OrderInfo
	)

	err = json.Unmarshal(respBs, &orderInfo)
	if err != nil {
		return nil, err
	}

	if orderInfo == nil {
		return nil, fmt.Errorf("get order is nil for remote")
	}
	return orderInfo, nil
}

// consume 确认订单[消耗型商品]
func (c *Client) consume(productId, purchaseToken string) error {
	// 检查刷新 access token
	if len(c.accessToken) <= 0 || c.accessTokenExpire.Unix() < time.Now().Unix() {
		// 刷新token
		_, err := c.getAccessToken()
		if err != nil {
			return err
		}
	}

	// 确认订单
	requestParams := consumeRequest{
		AccessToken: c.accessToken,
	}

	// new request
	req := ghttp.Request{
		Url:       fmt.Sprintf(consumeApi, c.packageName, productId, purchaseToken),
		Query:     requestParams,
		Method:    "POST",
		ShowDebug: client.debug,
		Timeout:   time.Duration(c.timeoutSecond) * time.Second,
	}
	req.AddHeader("Content-Type", "application/json")

	// send request
	resp, err := req.Do()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("counsume product http status code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

// acknowledge 确认订单[非消耗型商品]
func (c *Client) acknowledge(productId, purchaseToken string) error {
	// 检查刷新 access token
	if len(c.accessToken) <= 0 || c.accessTokenExpire.Unix() < time.Now().Unix() {
		// 刷新token
		_, err := c.getAccessToken()
		if err != nil {
			return err
		}
	}

	// 确认订单
	requestParams := consumeRequest{
		AccessToken: c.accessToken,
	}

	// new request
	req := ghttp.Request{
		Url:       fmt.Sprintf(acknowledgeApi, c.packageName, productId, purchaseToken),
		Body:      requestParams,
		Method:    "POST",
		ShowDebug: client.debug,
		Timeout:   time.Duration(c.timeoutSecond) * time.Second,
	}
	req.AddHeader("Content-Type", "application/json")

	// send request
	resp, err := req.Do()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("acknowledge product http status code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}
