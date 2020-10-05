package main

import (
	"math/rand"
	"net/url"
	"runtime"
	"time"

	"github.com/RouterScript/ProxyClient"
	"github.com/icrowley/fake"

	"github.com/corpix/peephole/pkg/log"
)

var (
	l       log.Logger
	addrs   = []string{"socks5://127.0.0.1:1338"}
	workers = 5
)

func init() {
	rand.Seed(1337)
	runtime.GOMAXPROCS(workers)

	l, _ = log.Create(log.Config{})
}

func client(addr string) proxyclient.Dial {
	var (
		n   = 5
		err error
	)

	for n > 0 {
		n--

		proxy, err := url.Parse(addr)
		if err != nil {
			continue
		}
		dial, err := proxyclient.NewClient(proxy)
		if err != nil {
			continue
		}

		return dial
	}

	l.Fatal().Err(err)
	return nil
}

func worker(dial proxyclient.Dial) {
	var (
		ports = []string{"80", "443", "5222"}
		ip    string
	)

	for {
		if rand.Intn(100) > 98 {
			ip = "[" + fake.IPv6() + "]"
		} else {
			ip = fake.IPv4()
		}

		addr := ip + ":" + ports[rand.Intn(len(ports))]
		conn, err := dial("tcp", addr)
		if err != nil {
			l.Error().Str("addr", addr).Err(err)
		} else {
			l.Info().Str("addr", addr).Msg("wow connected! closing.")
			conn.Close()
		}

		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
	}
}

func main() {
	for _, addr := range addrs {
		w := workers
		for w != 0 {
			go worker(client(addr))
			w--
		}
	}

	select {}
}
