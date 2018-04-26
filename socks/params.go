package socks

import (
	"context"
	"net"

	"github.com/corpix/loggers"
)

// Params is used to setup and configure a Server.
type Params struct {
	// Logger to use for this server.
	Logger loggers.Logger

	// Authenticators can be provided to implement custom authentication.
	// By default, "auth-less" mode is enabled.
	// For password-based auth use UserPassAuthenticator.
	Authenticators []Authenticator

	// Resolver can be provided to do custom name resolution.
	// Defaults to DNSResolver if not provided.
	Resolver Resolver

	// Rewriter can be used to transparently rewrite addresses.
	// This is invoked before the Rule is invoked.
	Rewriter Rewriter

	// Rule is provided to enable custom logic around permitting
	// various commands, etc. If not provided then PermitAll will be used.
	Rule Rule

	// Optional function for dialing out.
	Dial func(ctx context.Context, network, addr string) (net.Conn, error)
}

func ParamsWithDefaults(p Params) Params {
	if len(p.Authenticators) == 0 {
		p.Authenticators = []Authenticator{&NoAuthAuthenticator{}}
	}

	if p.Resolver == nil {
		p.Resolver = DNSResolver{}
	}

	if p.Rule == nil {
		p.Rule = PermitAll()
	}

	if p.Dial == nil {
		p.Dial = func(ctx context.Context, n, a string) (net.Conn, error) {
			return net.Dial(n, a)
		}
	}

	return p

}

func NewParams(l loggers.Logger) Params {
	return Params{Logger: l}
}