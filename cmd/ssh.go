package cmd

import (
	"os"
	"os/exec"

	"github.com/yamayo/ec2ssh/ec2"
)

func Run(instance *ec2.Instance, config SSHConfig) error {
	target := []string{
		"-i", config.IdentifyFile,
		config.User + "@" + instance.PublicIp + ":" + config.Port,
	}

	cmd := exec.Command("ssh", target, config.Options...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
