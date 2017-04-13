package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/yamayo/ec2ssh/ec2"
	"github.com/yamayo/ec2ssh/peco"
)

const ServerAliveInterval = 200

var (
	profile       *string = flag.String("profile", "default", "Use a specific profile from your credential file.")
	region        *string = flag.String("region", "ap-northeast-1", "The region to use. Overrides AWS config/env settings.")
	user          *string = flag.String("user", "ec2-user", "Login user name.")
	key           *string = flag.String("i", "", "Selects a file from which the identity (private key) for public key authentication is read.")
	port          *string = flag.String("port", "", "Port")
	aliveInterval *int    = flag.Int("interval", ServerAliveInterval, "")
)

func main() {
	flag.Parse()

	ec2 := ec2.NewClient(*profile, *region)
	instances, err := ec2.GetRunInstances()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	w := new(tabwriter.Writer)
	buffer := &bytes.Buffer{}
	w.Init(buffer, 4, 4, 4, '\t', 0)
	for _, inst := range instances {
		v := reflect.ValueOf(*inst)
		values := make([]string, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			values[i] = v.Field(i).String()
		}

		fmt.Fprintln(w, strings.Join(values, "\t"))
	}
	w.Flush()

	pec := peco.NewRunner()
	selected, err := pec.Transform(buffer.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	if len(selected) == 0 {
		os.Exit(0)
	}

	// words := strings.Fields(selected)
	// ip := words[1]
	// if len(*key) == 0 {
	// 	*key = "~/.ssh/" + words[5] + ".pem"
	// }

	// // config := &config.SSHConfig{User: *user, Port: *port, IdentityFile: *key}
	// // option := "ServerAliveInterval=" + strconv.Itoa(*aliveInterval)
	// // ssh.Run(instance, config)
	// target := []string{
	// 	"-i", *key,
	// 	*user + "@" + words[2] + ":" + *port,
	// }

	// cmd := exec.Command("ssh", target)
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// cmd.Run()
}
