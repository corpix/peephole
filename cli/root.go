package cli

import (
	"context"
	builtinLogger "log"
	"net"
	"time"

	//metrics "github.com/armon/go-metrics"
	socks "github.com/armon/go-socks5"
	"github.com/corpix/effects/writer"
	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"
	"github.com/urfave/cli"
)

var (
	// RootCommands is a list of subcommands for the application.
	RootCommands = []cli.Command{}

	// RootFlags is a list of flags for the application.
	RootFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "application configuration file",
			EnvVar: "CONFIG",
			Value:  "config.json",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "add this flag to enable debug mode",
		},
		cli.BoolFlag{
			Name:  "profile",
			Usage: "write profile information for debugging(cpu.prof, heap.prof)",
		},
		cli.BoolFlag{
			Name:  "trace",
			Usage: "write trace information for debugging(trace.prof)",
		},
	}
)

type ipNetPair struct {
	ip  net.IP
	net *net.IPNet
}

type ProxyAccess struct {
	log     loggers.Logger
	targets []ipNetPair
}

func (p *ProxyAccess) Allow(ctx context.Context, req *socks.Request) (context.Context, bool) {
	var (
		res = false
	)

	for _, target := range p.targets {
		if target.ip.Equal(req.DestAddr.IP) || target.net.Contains(req.DestAddr.IP) {
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

// RootAction is executing when program called without any subcommand.
func RootAction(c *cli.Context) error {
	var (
		targets = make([]ipNetPair, len(Config.Targets))
		cfg     *socks.Config
		err     error
	)

	cfg = &socks.Config{
		Logger: builtinLogger.New(
			writer.NewTrimSuffixWriter(log, []byte{'\n'}),
			"",
			0,
		),
		Resolver: socks.DNSResolver{},
	}

	if len(Config.Accounts) == 0 && len(Config.Targets) == 0 {
		log.Fatal("Neither Accounts or Targets was specified")
	}

	if len(Config.Accounts) > 0 {
		log.Printf(
			"Will use authentication, has %d accounts",
			len(Config.Accounts),
		)
		cfg.AuthMethods = []socks.Authenticator{
			socks.UserPassAuthenticator{
				Credentials: socks.StaticCredentials(Config.Accounts),
			},
		}
	} else {
		log.Print("Will NOT use authentication, has no accounts")
	}

	if len(Config.Targets) > 0 {
		for k, v := range Config.Targets {
			ip, net, err := net.ParseCIDR(v)
			if err != nil {
				log.Fatal(err)
			}
			targets[k] = ipNetPair{ip, net}

		}
		log.Printf(
			"Will use endpoint whitelists, has %d targets",
			len(Config.Targets),
		)
	} else {
		log.Print("Will NOT use endpoint whitelists, has no targets")
	}
	cfg.Rules = &ProxyAccess{
		log:     prefixwrapper.New("ProxyAccess: ", log),
		targets: targets,
	}

	server, err := socks.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = server.ListenAndServe("tcp", Config.Addr)
		if err != nil {
			log.Error(err)
		}

		time.Sleep(5 * time.Second)
	}
}
