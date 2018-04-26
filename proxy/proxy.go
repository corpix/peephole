package proxy

import (
	"errors"
	"net"
	"regexp"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/corpix/peephole/config"
	"github.com/corpix/peephole/socks"
)

func NewParams(c config.Config, l loggers.Logger) (socks.Params, error) {
	var (
		addresses = make([]IPNet, len(c.Addresses))
		domains   = make([]*regexp.Regexp, len(c.Domains))
		p         = socks.NewParams(l)

		ipNet IPNet
		r     *regexp.Regexp
		err   error
	)

	if len(c.Accounts) == 0 && len(c.Addresses) == 0 {
		// FIXME: error type
		return p, errors.New("Neither Accounts or Addresses was specified")
	}

	if len(c.Accounts) > 0 {
		l.Printf(
			"Will use authentication, have %d accounts",
			len(c.Accounts),
		)
		p.Authenticators = []socks.Authenticator{
			socks.UserPassAuthenticator{Credentials: socks.StaticCredentials(c.Accounts)},
		}
	} else {
		l.Print("Will NOT use authentication, has no accounts")
	}

	if len(c.Addresses) > 0 {
		for k, v := range c.Addresses {
			ipNet = IPNet{}

			ipNet.IP, ipNet.Net, err = net.ParseCIDR(v)
			if err != nil {
				return p, err
			}
			addresses[k] = ipNet
		}

		l.Printf(
			"Will use addresses whitelists, have %d addresses",
			len(c.Addresses),
		)
	} else {
		l.Print("Will NOT use addresses whitelists, has no addresses")
	}

	if len(c.Domains) > 0 {
		for k, v := range c.Domains {
			r, err = regexp.Compile(v)
			if err != nil {
				return p, err
			}

			domains[k] = r
		}

		l.Printf(
			"Will use domain whitelists, have %d domains",
			len(c.Domains),
		)
	} else {
		l.Print("Will NOT use domain whitelists, has no domains")
	}

	p.Rule = &Access{
		addresses: addresses,
		domains:   domains,
		log:       prefixwrapper.New("Access: ", l),
	}

	return p, nil
}
