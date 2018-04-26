package socks

import (
	"bytes"
	"testing"
)

func TestNoAuth(t *testing.T) {
	req := bytes.NewBuffer(nil)
	req.Write([]byte{1, NoAuth})

	s := New(NewParams(logger(t)))

	var resp bytes.Buffer
	ctx, err := s.authenticate(&resp, req)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if ctx.Method != NoAuth {
		t.Fatal("Invalid Context Method")
	}

	out := resp.Bytes()
	if !bytes.Equal(out, []byte{socks5Version, NoAuth}) {
		t.Fatalf("bad: %v", out)
	}
}

func TestPasswordAuth_Valid(t *testing.T) {
	req := bytes.NewBuffer(nil)
	req.Write([]byte{2, NoAuth, UserPassAuth})
	req.Write([]byte{1, 3, 'f', 'o', 'o', 3, 'b', 'a', 'r'})

	cred := StaticCredentials{"foo": "bar"}

	p := NewParams(logger(t))
	p.Authenticators = []Authenticator{UserPassAuthenticator{Credentials: cred}}

	s := New(p)

	var resp bytes.Buffer
	ctx, err := s.authenticate(&resp, req)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if ctx.Method != UserPassAuth {
		t.Fatal("Invalid Context Method")
	}

	val, ok := ctx.Payload["Username"]
	if !ok {
		t.Fatal("Missing key Username in auth context's payload")
	}

	if val != "foo" {
		t.Fatal("Invalid Username in auth context's payload")
	}

	out := resp.Bytes()
	if !bytes.Equal(out, []byte{socks5Version, UserPassAuth, 1, authSuccess}) {
		t.Fatalf("bad: %v", out)
	}
}

func TestPasswordAuth_Invalid(t *testing.T) {
	req := bytes.NewBuffer(nil)
	req.Write([]byte{2, NoAuth, UserPassAuth})
	req.Write([]byte{1, 3, 'f', 'o', 'o', 3, 'b', 'a', 'z'})

	cred := StaticCredentials{"foo": "bar"}
	p := NewParams(logger(t))
	p.Authenticators = []Authenticator{UserPassAuthenticator{Credentials: cred}}

	s := New(p)

	var resp bytes.Buffer
	ctx, err := s.authenticate(&resp, req)
	if err != UserAuthFailed {
		t.Fatalf("err: %v", err)
	}

	if ctx != nil {
		t.Fatal("Invalid Context Method")
	}

	out := resp.Bytes()
	if !bytes.Equal(out, []byte{socks5Version, UserPassAuth, 1, authFailure}) {
		t.Fatalf("bad: %v", out)
	}
}

func TestNoSupportedAuth(t *testing.T) {
	req := bytes.NewBuffer(nil)
	req.Write([]byte{1, NoAuth})

	cred := StaticCredentials{"foo": "bar"}
	p := NewParams(logger(t))
	p.Authenticators = []Authenticator{UserPassAuthenticator{Credentials: cred}}

	s := New(p)

	var resp bytes.Buffer
	ctx, err := s.authenticate(&resp, req)
	if err != NoSupportedAuth {
		t.Fatalf("err: %v", err)
	}

	if ctx != nil {
		t.Fatal("Invalid Context Method")
	}

	out := resp.Bytes()
	if !bytes.Equal(out, []byte{socks5Version, noAcceptable}) {
		t.Fatalf("bad: %v", out)
	}
}
