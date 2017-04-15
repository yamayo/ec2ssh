package ec2

import (
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Client struct {
	*ec2.EC2
}

func NewClient(profile, region string) *Client {
	cred := credentials.NewSharedCredentials("", profile)
	cfg := &aws.Config{
		Credentials: cred,
	}
	if len(region) > 0 {
		cfg.Region = aws.String(region)
	}
	sess, _ := session.NewSession(cfg)

	// svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	return &Client{svc}
}

type Instance struct {
	Name         string
	InstanceId   string
	InstanceType string
	PublicIp     string
	PrivateIp    string
	KeyName      string
}

func (c *Client) GetRunningInstances() ([]*Instance, error) {
	req := &ec2.DescribeInstancesInput{
		Filters: runningFilter(),
	}
	resp, err := c.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	var instances []*Instance
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			publicIp := "-"
			if inst.PublicIpAddress != nil {
				publicIp = *inst.PublicIpAddress
			}

			instance := &Instance{
				Name:         getName(inst),
				InstanceId:   *inst.InstanceId,
				InstanceType: *inst.InstanceType,
				PublicIp:     publicIp,
				PrivateIp:    *inst.PrivateIpAddress,
				KeyName:      *inst.KeyName,
			}
			instances = append(instances, instance)
		}
	}

	return instances, nil
}

func runningFilter() []*ec2.Filter {
	return []*ec2.Filter{
		&ec2.Filter{
			Name: aws.String("instance-state-name"),
			Values: []*string{
				aws.String("running"),
			},
		},
	}
}

func getName(inst *ec2.Instance) string {
	for _, t := range inst.Tags {
		if *t.Key == "Name" {
			return url.QueryEscape(*t.Value)
		}
	}

	if inst.PublicDnsName != nil {
		return *inst.PublicDnsName
	}

	return *inst.PrivateDnsName
}
