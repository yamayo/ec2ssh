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

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	return &Client{svc}
}

type Instance struct {
	Name         string
	PublicIp     string
	PrivateIp    string
	InstanceId   string
	InstanceType string
	KeyName      string
}

// func (inst *Instance) String() string {
// 	return fmt.Println(inst.Name, inst.PublicIp, inst.PrivateIp, inst.InstanceId, inst.InstanceType, inst.KeyName)
// }

func Convert(aa interface{}) []string {
	return aa.([]string)
}

func (c *Client) GetRunInstances() ([]*Instance, error) {
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
			var publicIp string
			if inst.PublicIpAddress != nil {
				publicIp = *inst.PublicIpAddress
			}

			instance := &Instance{
				Name:         getName(inst),
				PublicIp:     publicIp,
				PrivateIp:    *inst.PrivateIpAddress,
				InstanceId:   *inst.InstanceId,
				InstanceType: *inst.InstanceType,
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

	var name = ""
	if inst.PublicDnsName != nil {
		name = *inst.PublicDnsName
	} else if inst.PrivateDnsName != nil {
		name = *inst.PrivateDnsName
	}

	return name
}
