package proxy

import (
	"context"

	//metrics "github.com/armon/go-metrics"
	socks "github.com/armon/go-socks5"
	"github.com/corpix/loggers"
)

type Access struct {
	targets []IPNet
	log     loggers.Logger
}

func (p *Access) Allow(ctx context.Context, req *socks.Request) (context.Context, bool) {
	var (
		res = false
	)

	for _, target := range p.targets {
		if target.IP.Equal(req.DestAddr.IP) || target.Net.Contains(req.DestAddr.IP) {
			res = true
			break
		}
	}

	if res {
		p.log.Printf("Allow access %s -> %s", req.RemoteAddr.IP, req.DestAddr.IP)
	} else {
		p.log.Error("Deny access %s -> %s", req.RemoteAddr.IP, req.DestAddr.IP)
	}

	switch req.Command {
	case socks.ConnectCommand:
		return ctx, res
	}

	return ctx, false
}
