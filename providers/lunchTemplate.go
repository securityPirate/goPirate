package providers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//CreateLaunchTemplate on AWS
func (p *Provider) CreateLaunchTemplate() {
	svc := ec2.New(p.sess)
	input := &ec2.CreateLaunchTemplateInput{
		LaunchTemplateData: &ec2.RequestLaunchTemplateData{
			ImageId:      aws.String("ami-09558250a3419e7d0"), // ami can be dynamic
			InstanceType: aws.String("t3a.nano"), // type can be dynamic
			KeyName:      aws.String(*p.key.KeyName),
			NetworkInterfaces: []*ec2.LaunchTemplateInstanceNetworkInterfaceSpecificationRequest{
				{
					AssociatePublicIpAddress: aws.Bool(true),
					DeviceIndex:              aws.Int64(0),
					DeleteOnTermination:      aws.Bool(true),
				},
			},
			TagSpecifications: []*ec2.LaunchTemplateTagSpecificationRequest{
				{
					ResourceType: aws.String("instance"),
					Tags: []*ec2.Tag{
						{
							Key:   aws.String("Name"),
							Value: aws.String("gonhunt"),
						},
					},
				},
			},
		},
		LaunchTemplateName: aws.String("gonhunt-template"),
		VersionDescription: aws.String("Version1"),
	}

	result, err := svc.CreateLaunchTemplate(input)
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

	p.template = result
	fmt.Println(result)
}

//DeleteLaunchTemplate on AWS
func (p *Provider) DeleteLaunchTemplate() {
	svc := ec2.New(p.sess)
	input := &ec2.DeleteLaunchTemplateInput{
		LaunchTemplateId: aws.String(*p.template.LaunchTemplate.LaunchTemplateId),
	}

	result, err := svc.DeleteLaunchTemplate(input)
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
	p.template = nil
	fmt.Println(result)
}
