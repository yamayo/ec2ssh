package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yamayo/ec2ssh/ec2"
	"github.com/yamayo/ec2ssh/peco"
	"github.com/yamayo/ec2ssh/ssh"
	"github.com/yamayo/ec2ssh/util"
)

const version = "x.x.x"

var (
	profile *string = flag.String("profile", "", "Use a specific profile from your credential file.")
	region  *string = flag.String("region", "", "The region to use. Overrides AWS config/env settings.")
	user    *string = flag.String("user", "ec2-user", "Login user name.")
	keyPath *string = flag.String("identity file path", "~/.ssh", "Selects a file from which the identity (private key) for public key authentication is read.")
)

func main() {
	flag.Parse()

	ec2 := ec2.NewClient(*profile, *region)
	instances, err := ec2.GetRunningInstances()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	if len(instances) == 0 {
		fmt.Println("No running instances")
		os.Exit(0)
	}

	r := peco.NewRunner()
	selected, err := r.Select(util.Transform(instances))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	if len(selected) == 0 {
		os.Exit(0)
	}

	inst := util.RetriveInstance(selected, instances)
	cfg := ssh.ConfigFor(inst).WithUser(*user).WithIdentityFile(*keyPath, inst.KeyName)
	ssh.Run(cfg)
}
