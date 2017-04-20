package util

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/yamayo/ec2ssh/ec2"
)

func Transform(instances []*ec2.Instance) string {
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

	return buffer.String()
}

func RetrieveInstance(selected string, instances []*ec2.Instance) *ec2.Instance {
	instanceId := strings.Fields(selected)[1]
	for _, inst := range instances {
		if inst.InstanceId == instanceId {
			return inst
		}
	}

	return nil
}
