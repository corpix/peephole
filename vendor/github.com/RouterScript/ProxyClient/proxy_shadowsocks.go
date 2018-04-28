package proxyclient

import (
	"errors"
	"net"
	"net/url"

	ss "github.com/shadowsocks/go-shadowsocks2/core"
)

func newShadowsocksProxyClient(proxy *url.URL, upstreamDial Dial) (dial Dial, err error) {
	if proxy, err = decodedBase64EncodedURL(proxy); err != nil {
		return
	}
	if proxy.User == nil {
		err = errors.New("method and password is not available")
		return
	}
	var cipher ss.Cipher
	if password, ok := proxy.User.Password(); ok {
		method := proxy.User.Username()
		cipher, err = ss.PickCipher(method, nil, password)
		if err != nil {
			return
		}
	}
	dial = func(network, address string) (conn net.Conn, err error) {
		conn, err = upstreamDial(network, address)
		conn = cipher.StreamConn(conn)
		return
	}
	return
}
