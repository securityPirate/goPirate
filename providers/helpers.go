package providers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"net"
)

//ListIPs of aws instances
func (p Provider) ListIPs(counter int64) []net.IP {
	var ip []net.IP
	ec2svc := ec2.New(p.sess)
	inputAdd := &ec2.DescribeNetworkInterfacesInput{}
	resultAdd, err := ec2svc.DescribeNetworkInterfaces(inputAdd)
	if err != nil {
		fmt.Println("Error describe network", err)
	}
	for i := 0; i < int(counter); i++ {
		x := net.IP(*resultAdd.NetworkInterfaces[i].Association.PublicIp)
		ip = append(ip, x)
	}
	fmt.Println("the IPs has been fetched")
	return ip

}




//Instances all
func (p Provider) Instances() {
	svc := ec2.New(p.sess)
	input := &ec2.DescribeInstancesInput{}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//InstanceStaus all
func (p Provider) InstanceStaus() {
	svc := ec2.New(p.sess)
	input := &ec2.DescribeInstanceStatusInput{}

	result, err := svc.DescribeInstanceStatus(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
