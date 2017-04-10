package main

import (
	"flag"
	"fmt"
	"os"
	// "strconv"
	// "strings"

	// "github.com/yamayo/ec2ssh/config"
	"github.com/yamayo/ec2ssh/ec2"
	// "github.com/yamayo/ec2ssh/runner"
)

const ServerAliveInterval = 200

var (
	profile       *string = flag.String("profile", "default", "Use a specific profile from your credential file.")
	region        *string = flag.String("region", "ap-northeast-1", "The region to use. Overrides AWS config/env settings.")
	user          *string = flag.String("user", "ec2-user", "Login user name.")
	key           *string = flag.String("i", "", "Selects a file from which the identity (private key) for public key authentication is read.")
	aliveInterval *int    = flag.Int("interval", ServerAliveInterval, "")
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	ec2 := ec2.NewClient(*profile, *region)
	instances, err := ec2.GetRunInstances()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	for _, inst := range instances {
		fmt.Println("instances: ", inst)
	}
	// w := new(tabwriter.Writer)
	// buffer := &bytes.Buffer{}
	// w.Init(buffer, 4, 4, 4, '\t', 0)
	// // info := []string{"Name", "PublicIp", "PrivateIp", "InstanceId", "InstanceType", "KeyName"}
	// // fmt.Fprintln(w, strings.Join(info, "\t"))
	// for idx, _ := range resp.Reservations {
	// 	for _, inst := range resp.Reservations[idx].Instances {
	// 		var name string
	// 		for _, t := range inst.Tags {
	// 			if *t.Key == "Name" {
	// 				name = url.QueryEscape(*t.Value)
	// 				break
	// 			}
	// 		}

	// 		if name == "" && inst.PublicDnsName != nil {
	// 			name = *inst.PublicDnsName
	// 		}
	// 		if name == "" && inst.PrivateDnsName != nil {
	// 			name = *inst.PrivateDnsName
	// 		}
	// 		var publicIp string
	// 		if inst.PublicIpAddress != nil {
	// 			publicIp = *inst.PublicIpAddress
	// 		}

	// 		instance := []string{name, publicIp, *inst.PrivateIpAddress, *inst.InstanceId, *inst.InstanceType, *inst.KeyName}
	// 		fmt.Fprintln(w, strings.Join(instance, "\t"))
	// 	}
	// }
	// w.Flush()

	// pf := runner.NewRunner()
	// selected, err := pf.Transform(buffer.String())
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	// 	os.Exit(1)
	// }

	// if len(selected) == 0 {
	// 	os.Exit(0)
	// }

	// words := strings.Fields(selected)
	// ip := words[1]
	// if len(*key) == 0 {
	// 	*key = "~/.ssh/" + words[5] + ".pem"
	// }

	// config := &config.SSHConfig{User: *user, Port: *port, IdentityFile: *key}
	// option := "ServerAliveInterval=" + strconv.Itoa(*aliveInterval)
	// ssh.Run(instance, config)
}
