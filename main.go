package main

import (
	"flag"
	"os"
	"os/exec"
	"fmt"
	"github.com/yamayo/ec2ssh/runner"
	"bytes"
	"net/url"
	"strings"
	"strconv"
	"text/tabwriter"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const ServerAliveInterval = 200

var (
	region *string = flag.String("region", "ap-northeast-1", "The region to use. Overrides AWS config/env settings.")
	user *string = flag.String("user", "ec2-user", "Login user name.")
	key *string = flag.String("i", "", "Selects a file from which the identity (private key) for public key authentication is read.")
	profile *string = flag.String("profile", "default", "")
	aliveInterval *int = flag.Int("interval", ServerAliveInterval, "")
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	cred := credentials.NewSharedCredentials("", *profile)
	sess, err := session.NewSession(&aws.Config{
		Credentials: cred,
	})
	if err != nil {
	    panic(err)
	}
	svc := ec2.New(sess, &aws.Config{Region: aws.String(*region)})
	params := &ec2.DescribeInstancesInput{
        Filters: []*ec2.Filter{
            &ec2.Filter{
                Name: aws.String("instance-state-name"),
                Values: []*string{
                    aws.String("running"),
                },
            },
        },
    }
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		panic(err)
	}

	w := new(tabwriter.Writer)
	buffer := &bytes.Buffer{}
	w.Init(buffer, 4, 4, 4, '\t', 0)
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			var name string
            for _, t := range inst.Tags {
                if *t.Key == "Name" {
                    name = url.QueryEscape(*t.Value)
                }
            }

			if (name == "" && inst.PublicDnsName != nil) {
				name = *inst.PublicDnsName
			}
			if (name == "" && inst.PrivateDnsName != nil) {
				name = *inst.PrivateDnsName
			}
			var publicIp = ""
			if (inst.PublicIpAddress != nil) {
				publicIp = *inst.PublicIpAddress
			}

			instance := []string{name, publicIp, *inst.PrivateIpAddress, *inst.InstanceId, *inst.InstanceType, *inst.KeyName}
			fmt.Fprintln(w, strings.Join(instance, "\t"))
		}
	}
	w.Flush()

	pf := runner.NewRunner()
	selected, err := pf.Transform(buffer.String())
	if err != nil {
	  fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	  os.Exit(1)
	}

	if len(selected) == 0 {
		os.Exit(0)
	}

	words := strings.Fields(selected)
	ip := words[1]
	if len(*key) == 0 {
		*key = "~/.ssh/" + words[5] + ".pem"
	}

	option := "ServerAliveInterval="+strconv.Itoa(*aliveInterval)
	cmd := exec.Command("ssh", "-i", *key, *user + "@" + ip, "-o", option)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
