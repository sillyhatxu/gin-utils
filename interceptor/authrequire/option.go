package authrequire

import "github.com/sillyhatxu/gin-utils/jwtutils"

type Config struct {
	JWTClient  *jwtutils.JWT
	TokenKey   string
	ContextKey string
	IsDebug    bool
	DebugInput func() interface{}
}

type Option func(*Config)

func JWTClient(JWTClient *jwtutils.JWT) Option {
	return func(c *Config) {
		c.JWTClient = JWTClient
	}
}

func IsDebug(IsDebug bool) Option {
	return func(c *Config) {
		c.IsDebug = IsDebug
	}
}

func DebugInput(DebugInput func() interface{}) Option {
	return func(c *Config) {
		c.DebugInput = DebugInput
	}
}
