package ginutils

type Config struct {
	OpenHealth  bool
	ContextPath string
	SkipPaths   []string
}

type Option func(*Config)

func OpenHealth(openHealth bool) Option {
	return func(c *Config) {
		c.OpenHealth = openHealth
	}
}

func ContextPath(ContextPath string) Option {
	return func(c *Config) {
		c.ContextPath = ContextPath
	}
}

func SkipPaths(SkipPaths []string) Option {
	return func(c *Config) {
		c.SkipPaths = SkipPaths
	}
}
