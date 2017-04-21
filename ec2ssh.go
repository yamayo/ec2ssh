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
	profile     *string = flag.String("profile", "", `Use a specific profile from your credential file. (default "default")`)
	region      *string = flag.String("region", "", "The region to use. Overrides AWS config/env settings.")
	user        *string = flag.String("user", "ec2-user", "Specifies the user to login to EC2 machine.")
	showVersion *bool   = flag.Bool("version", false, "Show version")
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("ec2ssh version %s\n", version)
		os.Exit(0)
	}

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")

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

	inst := util.RetrieveInstance(selected, instances)

	cfg := ssh.NewConfig(inst).WithUser(*user)
	ssh.Run(cfg)
}
