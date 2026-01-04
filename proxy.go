package proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

var (
	dialer proxy.Dialer
)

// 取得proxy
//
// 使用VPN版本
func GetProxyDialerVPN(authUser, authPwd, proxyAddr, proxyPort string) (proxy.Dialer, error) {
	if dialer == nil {
		proxyAddr := fmt.Sprintf("%s:%s", proxyAddr, proxyPort)
		auth := proxy.Auth{User: authUser, Password: authPwd}

		var err error
		dialer, err = proxy.SOCKS5("tcp", proxyAddr, &auth, &net.Dialer{
			Timeout: 10 * time.Second,
		})
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrCreateSOCKS5DialerFailed, err)
		}
		logrus.Debugf("%s Proxy已成功設置", proxyAddr)
	}
	return dialer, nil
}

// 取得proxy
//
// 私有配置版本
func GetProxyDialer(apiURL, apiKey, proxyPort string) (proxy.Dialer, error) {
	if dialer == nil {
		// 1. 從 API 獲取當前 Proxy IP
		proxyIP, err := getProxyIP(apiURL, apiKey)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrGetProxyIPFailed, err)
		}
		logrus.Debugf("當前 Proxy IP: %s\n", proxyIP)

		// 2. 建立 SOCKS5 Dialer
		proxyAddr := fmt.Sprintf("%s:%s", proxyIP, proxyPort)

		// 3. 創建 SOCKS5 dialer
		dialer, err = proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrCreateSOCKS5DialerFailed, err)
		}
	}
	return dialer, nil
}

// 呼叫管理 API 獲取 IP
//
// 私有配置版本使用
func getProxyIP(apiURL, apiKey string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrCreateRequestFailed, err)
	}

	req.Header.Set("X-Allowed-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrHTTPRequestFailed, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%w: %d", ErrAPIStatusCodeError, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrReadResponseFailed, err)
	}

	return string(body), nil
}
