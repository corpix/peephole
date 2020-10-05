package proxy

import (
	"net"
	"regexp"

	"github.com/corpix/peephole/pkg/log"
	"github.com/corpix/peephole/pkg/proxy/metrics"
	"github.com/corpix/peephole/pkg/socks"
)

type Params = socks.Params

func CreateParams(c Config, m *metrics.Metrics, l log.Logger) (Params, error) {
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
		l.Info().Msgf(
			"will use authentication, have %d accounts",
			len(c.Accounts),
		)
		p.Authenticators = []socks.Authenticator{
			socks.UserPassAuthenticator{
				Credentials: socks.StaticCredentials(c.Accounts),
			},
		}
	} else {
		l.Warn().Msg("will NOT use authentication, has no accounts")
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

		l.Info().Msgf(
			"will use addresses whitelists, have %d addresses",
			len(c.Whitelist.Addresses),
		)
	} else {
		l.Warn().Msg("will NOT use addresses whitelists, has no addresses")
	}

	if len(c.Whitelist.Domains) > 0 {
		for k, v := range c.Whitelist.Domains {
			matcher, err = regexp.Compile(v)
			if err != nil {
				return p, err
			}

			domains[k] = matcher
		}

		l.Info().Msgf(
			"will use domain whitelists, have %d domains",
			len(c.Whitelist.Domains),
		)
	} else {
		l.Warn().Msg("will NOT use domain whitelists, has no domains")
	}

	//

	p.Metrics = m
	p.Rule = &Access{
		addresses: addresses,
		domains:   domains,
		log:       l,
	}

	//

	p.ReadDeadlineDuration = c.ReadDeadline
	p.WriteDeadlineDuration = c.WriteDeadline

	return p, nil
}
