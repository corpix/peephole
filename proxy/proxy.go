package proxy

import (
	"errors"
	"net"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/corpix/peephole/config"
	"github.com/corpix/peephole/socks"
)

func NewParams(c config.Config, l loggers.Logger) (socks.Params, error) {
	var (
		targets = make([]IPNet, len(c.Targets))
		p       = socks.NewParams(l)

		ipNet IPNet
		err   error
	)

	if len(c.Accounts) == 0 && len(c.Targets) == 0 {
		// FIXME: error type
		return p, errors.New("Neither Accounts or Targets was specified")
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

	if len(c.Targets) > 0 {
		for k, v := range c.Targets {
			ipNet = IPNet{}

			ipNet.IP, ipNet.Net, err = net.ParseCIDR(v)
			if err != nil {
				return p, err
			}
			targets[k] = ipNet
		}

		l.Printf(
			"Will use endpoint whitelists, have %d targets",
			len(c.Targets),
		)
	} else {
		l.Print("Will NOT use endpoint whitelists, has no targets")
	}

	p.Rule = &Access{
		targets: targets,
		log:     prefixwrapper.New("Access: ", l),
	}

	return p, nil
}
