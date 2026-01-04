package proxy

import (
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

var (
	dialer proxy.Dialer
)

// 取得proxy dialer
func GetProxyDialer(proxyAddr string, proxyAuth *proxy.Auth, proxyPort string) (proxy.Dialer, error) {
	if dialer == nil {
		proxyAddr := fmt.Sprintf("%s:%s", proxyAddr, proxyPort)

		var err error
		if proxyAuth != nil {
			dialer, err = proxy.SOCKS5("tcp", proxyAddr, proxyAuth, &net.Dialer{
				Timeout: 10 * time.Second,
			})
		} else {
			dialer, err = proxy.SOCKS5("tcp", proxyAddr, nil, &net.Dialer{
				Timeout: 10 * time.Second,
			})
		}

		if err != nil {
			dialer = nil
			return nil, fmt.Errorf("%w: %v", ErrCreateSOCKS5DialerFailed, err)
		}
		logrus.Debugf("%s Proxy已成功設置", proxyAddr)
	}
	return dialer, nil
}

func GenerateProxyAuth(authUser, authPwd string) *proxy.Auth {
	return &proxy.Auth{User: authUser, Password: authPwd}
}
