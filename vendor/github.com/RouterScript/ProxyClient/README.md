# ProxyClient

the proxy client library

supported SOCKS4, SOCKS4A, SOCKS5, HTTP, HTTPS etc proxy protocols

## Supported Schemes
- [x] Direct
- [x] Reject
- [x] Blackhole
- [x] HTTP
- [x] HTTPS
- [x] SOCKS4
- [x] SOCKS4A
- [x] SOCKS5
- [x] SOCKS5 with TLS
- [x] ShadowSocks
- [x] SSH Agent
- [ ] VMess

# Documentation

The full documentation is available on [Godoc](//godoc.org/github.com/RouterScript/ProxyClient).

# Example
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/RouterScript/ProxyClient"
)

func main() {
	proxy, _ := url.Parse("http://localhost:8080")
	dial, _ := proxyclient.NewClient(proxy)
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: dial.Context,
		},
	}
	request, err := client.Get("http://www.example.com")
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
```

# Reference

see http://github.com/GameXG/ProxyClient
