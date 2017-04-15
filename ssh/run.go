package ssh

import (
	"os"
	"os/exec"
)

func Run(cfg *Config) {
	opts := []string{
		"-i", cfg.IdentityFile,
		*cfg.User + "@" + cfg.Ip,
	}
	cmd := exec.Command("ssh", opts...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
