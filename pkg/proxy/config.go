package proxy

import (
	"time"

	"github.com/corpix/peephole/pkg/proxy/metrics"
)

type Config struct {
	Accounts  map[string]string
	Whitelist WhitelistConfig
	Metrics   *metrics.Config

	ReadDeadline  time.Duration
	WriteDeadline time.Duration
}

func (c *Config) Default() {
loop:
	for {
		switch {
		case c.Metrics == nil:
			c.Metrics = &metrics.Config{}
		case c.ReadDeadline == 0:
			c.ReadDeadline = 60 * time.Second
		case c.WriteDeadline == 0:
			c.WriteDeadline = 60 * time.Second
		default:
			break loop
		}
	}
}

//

type WhitelistConfig struct {
	Addresses []string
	Domains   []string
}
