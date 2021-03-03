package providers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"gopirate.com/local"
	"log"
	"time"
)

//Provider AWS
type Provider struct {
	name     string
	sess     *session.Session
	key      *ec2.CreateKeyPairOutput
	template *ec2.CreateLaunchTemplateOutput
	fleet    *ec2.RequestSpotFleetOutput
}

type key struct {
	name string
	path string
}

//ConnectAndLoad with aws for testing
func (p *Provider) ConnectAndLoad(name, access, secret, region string, logger *log.Logger , config []byte) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region), //region can be dynamic like access and secret :D
		Credentials: credentials.NewStaticCredentials(access, secret, ""),
	})
	if err != nil {
		logger.Println("Error creating session ", err)
	}
	p.name = name
	p.sess = sess
	logger.Println("Session Created on", name)
}

//ConnectAndLunch with aws for testing
func (p *Provider) ConnectAndLunch(name, access, secret, region string, counter int64, logger *log.Logger , config []byte) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region), //region can be dynamic like access and secret :D
		Credentials: credentials.NewStaticCredentials(access, secret, ""),
	})
	if err != nil {
		logger.Println("Error creating session ", err)
	}
	p.name = name
	p.sess = sess
	logger.Println("Session Created on", name)
	//detect .goPirate dir
	//read the fleet data from there if not exist creat new fleet

	fmt.Println("Do you wanna to Create new Fleet? press y")

	if local.Answer() {
		p.CreateKey()
		p.CreateLaunchTemplate()
		p.CreateFleet(counter)
		fmt.Println("Please wait, We are preparing your fleet")
		time.Sleep(90 * time.Second)
	}

}

//Flush deleting everything
func (p Provider) Flush() {
	fmt.Print("Deleting....")
	p.DeleteKey()
	p.DeleteLaunchTemplate()
	p.DeleteFleet()
	fmt.Println(".................................................Done")
}
