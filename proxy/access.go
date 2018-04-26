package proxy

import (
	"context"

	"github.com/corpix/loggers"

	"github.com/corpix/peephole/socks"
)

type Access struct {
	targets []IPNet
	log     loggers.Logger
}

func (p *Access) Match(ctx context.Context, req *socks.Request) (context.Context, bool) {
	res := false

	for _, target := range p.targets {
		if target.IP.Equal(req.DestAddr.IP) || target.Net.Contains(req.DestAddr.IP) {
			res = true
			break
		}
	}

	if res {
		p.log.Printf("Allow '%s' to '%s'", req.RemoteAddr.IP, req.DestAddr.IP)
	} else {
		p.log.Error("Deny '%s' to '%s'", req.RemoteAddr.IP, req.DestAddr.IP)
	}

	switch req.Command {
	case socks.ConnectCommand:
		return ctx, res
	}

	return ctx, false
}
