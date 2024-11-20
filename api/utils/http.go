package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const maxRetries = 3
const retryDelay = 2 * time.Second

// Request 普通请求
func Request(ctx context.Context, method string, urls string, body io.Reader) (*http.Response, error) {
	// 创建请求函数
	createRequest := func(ctx context.Context) (*http.Request, error) {
		request, err := http.NewRequestWithContext(ctx, strings.ToUpper(method), urls, body)
		if err != nil {
			return nil, err
		}
		if strings.ToUpper(method) == "POST" {
			request.Header.Set("Content-Type", "application/json")
		}

		return request, nil
	}

	// 最大重试次数
	for i := 0; i < maxRetries; i++ {
		// 为每次请求创建一个新的 context，并设置超时
		reqCtx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		// 创建请求
		request, err := createRequest(reqCtx)
		if err != nil {
			cancelFunc()
			return nil, err
		}
		// 发送请求
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			cancelFunc()
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			return nil, fmt.Errorf("请求失败: %w", err)
		}
		cancelFunc()
		// 请求成功
		return response, nil
	}

	return nil, errors.New("请求失败，未知的错误！")
}

// BmRequest  巨量带token请求
func BmRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Response, error) {
	// 配置项只获取一次
	AccessToken := viper.GetString("AccessToken") // todo 从redis获取
	url := viper.GetString("Bm.Host") + path

	// 创建请求函数
	createRequest := func(ctx context.Context) (*http.Request, error) {
		request, err := http.NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			return nil, err
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("AccessToken", AccessToken)
		return request, nil
	}

	// 最大重试次数
	for i := 0; i < maxRetries; i++ {
		// 为每次请求创建一个新的 context，并设置超时
		reqCtx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
		// defer cancelFunc() // 确保每次都能取消

		// 创建请求
		request, err := createRequest(reqCtx)
		if err != nil {
			cancelFunc()
			return nil, err
		}

		// 发送请求
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			cancelFunc()
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			return nil, fmt.Errorf("请求失败: %w", err)
		}
		cancelFunc()
		// 请求成功
		return response, nil
	}

	// 达到最大重试次数仍未成功
	return nil, errors.New("请求失败，已达到最大重试次数")
}

func ParseUrl(urls string) error {

	uri, err := url.ParseRequestURI(urls)
	if err != nil {
		return err
	}

	if uri.Scheme != "http" && uri.Scheme != "https" {
		return errors.New("协议错误")
	}

	if uri.Host == "" {
		return errors.New("域名错误")
	}

	return nil
}
