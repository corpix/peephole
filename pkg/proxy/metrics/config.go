package metrics

type Config struct {
	ServiceName     string
	StatsdAddresses []string
}

func (c *Config) Default() {
loop:
	for {
		switch {
		case c.ServiceName == "":
			c.ServiceName = "peephole"
		default:
			break loop
		}
	}
}
