package ssh

import (
	"log"
	"os/user"
	"path"

	"github.com/yamayo/ec2ssh/ec2"
)

const ServerAliveInterval = "200"

type Config struct {
	Ip           string
	User         string
	IdentityFile string
}

func NewConfig(inst *ec2.Instance) *Config {
	cfg := &Config{
		Ip:           inst.Ip(),
		IdentityFile: privateKeyPath(inst.KeyName),
	}

	return cfg
}

func (c *Config) WithUser(user string) *Config {
	c.User = user
	return c
}

func privateKeyPath(keyName string) string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(user.HomeDir, ".ssh", keyName+".pem")
}
