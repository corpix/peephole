package proxyclient

import (
	"context"
	"errors"
	"net"
	"net/url"
	"strings"
)

type Dial func(network, address string) (net.Conn, error)

type DialFactory func(*url.URL, Dial) (Dial, error)

var schemes = map[string]DialFactory{
	"DIRECT":     newDirectProxyClient,
	"REJECT":     newRejectProxyClient,
	"BLACKHOLE":  newBlackholeProxyClient,
	"SOCKS":      newSocksProxyClient,
	"SOCKS4":     newSocksProxyClient,
	"SOCKS4A":    newSocksProxyClient,
	"SOCKS5":     newSocksProxyClient,
	"SOCKS5+TLS": newSocksProxyClient,
	"HTTP":       newHTTPProxyClient,
	"HTTPS":      newHTTPProxyClient,
	"SS":         newShadowsocksProxyClient,
	"SSH":        newSSHAgentProxyClient,
}

func NewClient(proxy *url.URL) (Dial, error) {
	return NewClientWithDial(proxy, net.Dial)
}

func NewClientChain(proxies []*url.URL) (Dial, error) {
	return NewClientChainWithDial(proxies, net.Dial)
}

func NewClientWithDial(proxy *url.URL, upstreamDial Dial) (_ Dial, err error) {
	if proxy == nil {
		err = errors.New("proxy url is nil")
		return
	}
	if upstreamDial == nil {
		err = errors.New("upstream dial is nil")
		return
	}
	proxy = normalizeLink(*proxy)
	if _, ok := schemes[proxy.Scheme]; !ok {
		err = errors.New("unsupported proxy client.")
		return
	}
	return schemes[proxy.Scheme](proxy, upstreamDial)
}

func NewClientChainWithDial(proxies []*url.URL, upstreamDial Dial) (dial Dial, err error) {
	dial = upstreamDial
	for _, proxyURL := range proxies {
		dial, err = NewClientWithDial(proxyURL, dial)
		if err != nil {
			return
		}
	}
	return
}

func RegisterScheme(schemeName string, factory DialFactory) {
	schemes[strings.ToUpper(schemeName)] = factory
}

func SupportedSchemes() []string {
	schemeNames := make([]string, 0, len(schemes))
	for schemeName := range schemes {
		schemeNames = append(schemeNames, schemeName)
	}
	return schemeNames
}

func (dial Dial) Context(ctx context.Context, network, address string) (net.Conn, error) {
	return dial(network, address)
}

func (dial Dial) TCPOnly(network, address string) (net.Conn, error) {
	switch strings.ToUpper(network) {
	case "TCP", "TCP4", "TCP6":
		return dial(network, address)
	default:
		return nil, errors.New("unsupported network type.")
	}
}

func (dial Dial) Dial(network, address string) (net.Conn, error) {
	return dial(network, address)
}
