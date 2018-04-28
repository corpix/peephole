package main

import (
	"math/rand"
	"net/url"
	"runtime"
	"time"

	"github.com/RouterScript/ProxyClient"
	"github.com/icrowley/fake"
	"github.com/sirupsen/logrus"
)

var (
	log     = logrus.New()
	addrs   = []string{"socks5://127.0.0.1:1338"}
	workers = 5
)

func init() {
	rand.Seed(1337)
	runtime.GOMAXPROCS(workers)
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

	log.Fatal(err)
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
			log.Error(addr, " ", err)
		} else {
			log.Print(addr, " wow connected! Closing!")
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
