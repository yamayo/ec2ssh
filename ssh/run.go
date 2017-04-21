package ssh

import (
	"os"
	"os/exec"
)

func Run(cfg *Config) {
	args := []string{
		"-i", cfg.IdentityFile,
		cfg.User + "@" + cfg.Ip,
		"-o", "ServerAliveInterval=" + ServerAliveInterval,
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
