package proxy

import (
	"errors"
	"net"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"

	"github.com/corpix/peephole/config"
	"github.com/corpix/peephole/socks"
)

func NewConfig(c config.Config, l loggers.Logger) (*socks.Config, error) {
	var (
		targets = make([]IPNet, len(c.Targets))
		cfg     = socks.Config{}

		ipNet IPNet
		err   error
	)

	if len(c.Accounts) == 0 && len(c.Targets) == 0 {
		return nil, errors.New("Neither Accounts or Targets was specified")
	}

	if len(c.Accounts) > 0 {
		l.Printf(
			"Will use authentication, has %d accounts",
			len(c.Accounts),
		)
		cfg.AuthMethods = []socks.Authenticator{
			socks.UserPassAuthenticator{
				Credentials: socks.StaticCredentials(c.Accounts),
			},
		}
	} else {
		l.Print("Will NOT use authentication, has no accounts")
	}

	if len(c.Targets) > 0 {
		for k, v := range c.Targets {
			ipNet = IPNet{}

			ipNet.IP, ipNet.Net, err = net.ParseCIDR(v)
			if err != nil {
				l.Fatal(err)
			}
			targets[k] = ipNet
		}

		l.Printf(
			"Will use endpoint whitelists, has %d targets",
			len(c.Targets),
		)
	} else {
		l.Print("Will NOT use endpoint whitelists, has no targets")
	}

	cfg.Rules = &Access{
		log:     prefixwrapper.New("Proxy Access: ", l),
		targets: targets,
	}

	return &cfg, nil
}
