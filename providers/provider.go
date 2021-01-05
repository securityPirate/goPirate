package providers

import (
	"time"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//Provider AWS
type Provider struct {
	name     string
	sess     *session.Session
	key      *ec2.CreateKeyPairOutput
	template *ec2.CreateLaunchTemplateOutput
	fleet    *ec2.RequestSpotFleetOutput
}

//ConnectAndLunch with aws for testing
func (p *Provider) ConnectAndLunch(name, access, secret, region string, counter int64) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region), //region can be dynamic like access and secret :D
		Credentials: credentials.NewStaticCredentials(access, secret, ""),
	})
	if err != nil {
		fmt.Println("Error creating session ", err)
	}
	p.name = name
	p.sess = sess
	fmt.Println("Session Created")
	fmt.Println("Do you wanna to Create new Fleet? press y")
	var input string
	fmt.Scan(&input)
	if input == "y" {
		p.CreateKey()
		p.CreateLaunchTemplate()
		p.CreateFleet(counter)
		time.Sleep(3 * time.Minute)
	}

}

//Flush deleting everything
func (p Provider) Flush() {
	p.DeleteKey()
	p.DeleteLaunchTemplate()
	p.DeleteFleet()
}
