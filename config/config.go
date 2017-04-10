package config

type SSHConfig struct {
	User         string
	Port         int
	IdentityFile string
	Options      []string
}
