package proxy

import (
	"context"
	"regexp"

	"github.com/corpix/peephole/pkg/log"
	"github.com/corpix/peephole/pkg/socks"
)

type Access struct {
	addresses []IPNet
	domains   []*regexp.Regexp
	log       log.Logger
}

func (p *Access) Match(ctx context.Context, req *socks.Request) (context.Context, bool) {
	res := len(p.addresses) == 0 && len(p.domains) == 0

	if !res && req.DestAddr.FQDN != "" {
		for _, domain := range p.domains {
			if domain.MatchString(req.DestAddr.FQDN) {

			}
		}
	}

	if !res {
		for _, address := range p.addresses {
			if address.IP.Equal(req.DestAddr.IP) || address.Net.Contains(req.DestAddr.IP) {
				res = true
				break
			}
		}
	}

	if res {
		p.log.Debug().Msgf("Allow '%s' to '%s'", req.RemoteAddr, req.DestAddr)
	} else {
		p.log.Warn().Msgf("Deny '%s' to '%s'", req.RemoteAddr, req.DestAddr)
	}

	switch req.Command {
	case socks.ConnectCommand:
		return ctx, res
	}

	return ctx, false
}
