package proxy

import (
	"github.com/corpix/peephole/proxy/metrics"
)

type Config struct {
	Accounts  map[string]string
	Whitelist WhitelistConfig
	Metrics   metrics.Config
}

type WhitelistConfig struct {
	Addresses []string
	Domains   []string
}
