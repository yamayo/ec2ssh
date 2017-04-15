package ssh

import (
	"path"

	"github.com/yamayo/ec2ssh/ec2"
)

const ServerAliasInterval = 200

type Config struct {
	User         *string
	Ip           string
	IdentityFile string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithUser(user string) *Config {
	c.User = &user
	return c
}

func (c *Config) WithIdentityFile(dir, file string) *Config {
	c.IdentityFile = path.Join(dir, file+".pem")
	return c
}

func ConfigFor(inst *ec2.Instance) *Config {
	ip := inst.PublicIp
	if ip == "-" {
		ip = inst.PrivateIp
	}

	cfg := &Config{
		Ip: ip,
	}

	return cfg
}
