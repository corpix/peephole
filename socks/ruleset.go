package socks

import (
	"context"
)

// Rule is used to provide custom rules to allow or prohibit actions
type Rule interface {
	Match(ctx context.Context, req *Request) (context.Context, bool)
}

// PermitAll returns a Rule which allows all types of connections
func PermitAll() Rule {
	return &PermitCommand{true, true, true}
}

// PermitNone returns a Rule which disallows all types of connections
func PermitNone() Rule {
	return &PermitCommand{false, false, false}
}

// PermitCommand is an implementation of the Rule which
// enables filtering supported commands
type PermitCommand struct {
	EnableConnect   bool
	EnableBind      bool
	EnableAssociate bool
}

func (p *PermitCommand) Match(ctx context.Context, req *Request) (context.Context, bool) {
	switch req.Command {
	case ConnectCommand:
		return ctx, p.EnableConnect
	case BindCommand:
		return ctx, p.EnableBind
	case AssociateCommand:
		return ctx, p.EnableAssociate
	}

	return ctx, false
}
