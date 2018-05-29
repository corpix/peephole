package metrics

type Config struct {
	ServiceName     string   `validate:"required" default:"peephole"`
	StatsdAddresses []string `env:"STATSD_ADDRESSES"`
}
