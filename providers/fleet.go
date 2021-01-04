package providers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//CreateFleet on AWS
func (p *Provider) CreateFleet(counter int64) {
	svc := ec2.New(p.sess)
	input := &ec2.RequestSpotFleetInput{
		SpotFleetRequestConfig: &ec2.SpotFleetRequestConfigData{
			IamFleetRole: aws.String("arn:aws:iam::012114012049:role/aws-ec2-spot-fleet-tagging-role"), //create iamfleetrole dynamicly
			LaunchTemplateConfigs: []*ec2.LaunchTemplateConfig{
				{
					LaunchTemplateSpecification: &ec2.FleetLaunchTemplateSpecification{
						LaunchTemplateId: aws.String(*p.template.LaunchTemplate.LaunchTemplateId),
						Version:          aws.String("$Default"),
					},
				},
			},
			SpotPrice:      aws.String("0.01"),
			TargetCapacity: aws.Int64(counter),
			AllocationStrategy : aws.String("lowestPrice"),
		},
	}

	result, err := svc.RequestSpotFleet(input)
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
	p.fleet = result
	fmt.Println(result)
}

//DeleteFleet from AWS
func (p *Provider) DeleteFleet() {
	svc := ec2.New(p.sess)
	input := &ec2.CancelSpotFleetRequestsInput{
		SpotFleetRequestIds: []*string{
			aws.String(*p.fleet.SpotFleetRequestId),
		},
		TerminateInstances: aws.Bool(true),
	}

	result, err := svc.CancelSpotFleetRequests(input)
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
	p.fleet = nil
	fmt.Println(result)
}
