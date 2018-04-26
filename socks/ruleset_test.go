package socks

import (
	"context"
	"testing"
)

func TestPermitCommand(t *testing.T) {
	ctx := context.Background()
	r := &PermitCommand{true, false, false}

	if _, ok := r.Match(ctx, &Request{Command: ConnectCommand}); !ok {
		t.Fatalf("expect connect")
	}

	if _, ok := r.Match(ctx, &Request{Command: BindCommand}); ok {
		t.Fatalf("do not expect bind")
	}

	if _, ok := r.Match(ctx, &Request{Command: AssociateCommand}); ok {
		t.Fatalf("do not expect associate")
	}
}
