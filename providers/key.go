package providers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"io/ioutil"
)

//CreateKey on aws
func (p *Provider) CreateKey() error {
	svc := ec2.New(p.sess)
	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String("gonhuntKey"), //keyname can be dynamic
	}

	result, err := svc.CreateKeyPair(input)
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
		return err
	}
	p.key = result
	fmt.Println(result)

	bytes := []byte(*p.key.KeyMaterial)
	saveFileTo := "Get user home/.goPirate/" + *p.key.KeyName + ".pem"
	fmt.Println("saving to file ...")
	errFile := ioutil.WriteFile("/home/ahmed/.goPirate/"+*p.key.KeyName+".pem", bytes, 0600)
	if errFile != nil {
		return errFile
	}

	fmt.Printf("Key saved to: %s", saveFileTo)
	return nil
}

//DeleteKey from AWS
func (p *Provider) DeleteKey() {
	svc := ec2.New(p.sess)
	input := &ec2.DeleteKeyPairInput{
		KeyName: aws.String(*p.key.KeyName),
	}

	result, err := svc.DeleteKeyPair(input)
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
	p.key = nil
	fmt.Println(result)
}
