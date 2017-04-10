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
	sess, _ := session.NewSession(&aws.Config{
		Credentials: cred,
	})
	// if err != nil {
	// 	return nil, err
	// }

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	return &Client{svc}
}

type Instance struct {
	Name         string
	PublicIp     string
	PrivateIp    string
	ID           string
	InstanceType string
	KeyName      string
}

type Instances []Instance

func (c *Client) GetRunInstances() ([]Instance, error) {
	params := &ec2.DescribeInstancesInput{
		Filters: runFilter,
	}
	resp, err := c.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	var instances Instances
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			var publicIp string
			if inst.PublicIpAddress != nil {
				publicIp = *inst.PublicIpAddress
			}

			instance := Instance{getName(inst), publicIp, *inst.PrivateIpAddress, *inst.InstanceId, *inst.InstanceType, *inst.KeyName}
			instances = append(instances, instance)
		}
	}

	return instances, nil
}

func runFilter() Filter {
	return []*ec2.Filter{
		&ec2.Filter{
			Name: aws.String("instance-state-name"),
			Values: []*string{
				aws.String("running"),
			},
		},
	}
}

func getName(inst ec2.Instance) string {
	for _, t := range inst.Tags {
		if *t.Key == "Name" {
			return url.QueryEscape(*t.Value)
		}
	}

	if inst.PublicDnsName != nil {
		return *inst.PublicDnsName
	} else if inst.PrivateDnsName != nil {
		return *inst.PrivateDnsName
	}
}