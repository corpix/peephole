package proxy

type Config struct {
	Accounts  map[string]string
	StatsAddr string `env:"STATS_ADDR`
	Whitelist WhitelistConfig
}

type WhitelistConfig struct {
	Addresses []string
	Domains   []string
}
