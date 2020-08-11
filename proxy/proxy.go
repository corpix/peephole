package proxy

import (
	"net"
	"regexp"
	"time"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/corpix/peephole/proxy/metrics"
	"github.com/corpix/peephole/socks"
)

func NewParams(c Config, l loggers.Logger) (socks.Params, error) {
	var (
		p = socks.NewParams(l)

		addresses = make([]IPNet, len(c.Whitelist.Addresses))
		domains   = make([]*regexp.Regexp, len(c.Whitelist.Domains))

		ipNet   IPNet
		matcher *regexp.Regexp
		err     error
	)

	//

	if len(c.Accounts) > 0 {
		l.Printf(
			"Will use authentication, have %d accounts",
			len(c.Accounts),
		)
		p.Authenticators = []socks.Authenticator{
			socks.UserPassAuthenticator{
				Credentials: socks.StaticCredentials(c.Accounts),
			},
		}
	} else {
		l.Print("Will NOT use authentication, has no accounts")
	}

	//

	if len(c.Whitelist.Addresses) > 0 {
		for k, v := range c.Whitelist.Addresses {
			ipNet = IPNet{}

			ipNet.IP, ipNet.Net, err = net.ParseCIDR(v)
			if err != nil {
				return p, err
			}
			addresses[k] = ipNet
		}

		l.Printf(
			"Will use addresses whitelists, have %d addresses",
			len(c.Whitelist.Addresses),
		)
	} else {
		l.Print("Will NOT use addresses whitelists, has no addresses")
	}

	if len(c.Whitelist.Domains) > 0 {
		for k, v := range c.Whitelist.Domains {
			matcher, err = regexp.Compile(v)
			if err != nil {
				return p, err
			}

			domains[k] = matcher
		}

		l.Printf(
			"Will use domain whitelists, have %d domains",
			len(c.Whitelist.Domains),
		)
	} else {
		l.Print("Will NOT use domain whitelists, has no domains")
	}

	//

	m, err := metrics.New(c.Metrics, l)
	if err != nil {
		return p, err
	}

	p.Metrics = m

	//

	p.Rule = &Access{
		addresses: addresses,
		domains:   domains,
		log:       prefixwrapper.New("Access: ", l),
	}

	//

	p.ReadDeadlineDuration = 15 * time.Second
	p.WriteDeadlineDuration = 15 * time.Second

	return p, nil
}
