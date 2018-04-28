package proxy

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

type Client interface {
	Dial(network, address string) (net.Conn, error)
}

func NewClient(proxy *url.URL, conf *SOCKSConf) (client Client, err error) {
	switch strings.ToUpper(proxy.Scheme) {
	case "SOCKS4", "SOCKS4A":
		client = &socks4Client{proxy, conf}
	case "SOCKS5", "SOCKS5+TLS":
		client = &socks5Client{proxy, conf, conf.TLSConfig}
	default:
		err = fmt.Errorf("%s not supported", proxy.Scheme)
	}
	return
}

func splitHostPort(addr string) (host, port []byte, err error) {
	hostName, hostPort, err := net.SplitHostPort(addr)
	if err != nil {
		return
	}
	return []byte(hostName), toPort(hostPort), err
}
