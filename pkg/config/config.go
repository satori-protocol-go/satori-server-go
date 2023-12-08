package config

type Cfg interface {
	Token() string
	Addr() string
}

type simpleConfig struct {
	AddrString  string `yaml:"addr" json:"addr"`
	TokenString string `yaml:"token,omitempty" json:"token,omitempty"`
}

func (c *simpleConfig) Addr() string {
	if c.AddrString == "" {
		return ":8080"
	}
	return c.AddrString
}

func (c *simpleConfig) Token() string {
	return c.TokenString
}

func Load(path string) (Cfg, error) {
	return &simpleConfig{}, nil
}
