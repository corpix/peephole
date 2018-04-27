package proxy

type Config struct {
	Accounts  map[string]string
	Whitelist WhitelistConfig
	Metrics   MetricsConfig
}

type WhitelistConfig struct {
	Addresses []string
	Domains   []string
}

type MetricsConfig struct {
	ServiceName     string   `validate:"required" default:"peephole"`
	StatsdAddresses []string `env:"STATSD_ADDRESSES"`
}
