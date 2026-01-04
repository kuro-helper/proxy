package proxy

import "errors"

var (
	// ErrCreateRequestFailed 創建 HTTP 請求失敗
	ErrCreateRequestFailed = errors.New("proxy: failed to create HTTP request")
	// ErrHTTPRequestFailed HTTP 請求執行失敗
	ErrHTTPRequestFailed = errors.New("proxy: HTTP request execution failed")
	// ErrAPIStatusCodeError API 返回錯誤狀態碼
	ErrAPIStatusCodeError = errors.New("proxy: API returned error status code")
	// ErrReadResponseFailed 讀取響應體失敗
	ErrReadResponseFailed = errors.New("proxy: failed to read response body")
	// ErrGetProxyIPFailed 獲取 Proxy IP 失敗
	ErrGetProxyIPFailed = errors.New("proxy: failed to get proxy IP")
	// ErrCreateSOCKS5DialerFailed 建立 SOCKS5 Dialer 失敗
	ErrCreateSOCKS5DialerFailed = errors.New("proxy: failed to create SOCKS5 dialer")
	// ErrProxyRequestFailed 通過代理的 HTTP 請求失敗
	ErrProxyRequestFailed = errors.New("proxy: HTTP request through proxy failed")
)
