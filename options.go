package ginutils

type Config struct {
	OpenHealth bool
}

type Option func(*Config)

func OpenHealth(openHealth bool) Option {
	return func(c *Config) {
		c.OpenHealth = openHealth
	}
}
