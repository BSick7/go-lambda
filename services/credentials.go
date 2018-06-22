package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

func DefaultCredentials(conf *aws.Config) (*credentials.Credentials, error) {
	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, fmt.Errorf("error creating aws default credentials: %s", err)
	}

	return credentials.NewChainCredentials([]credentials.Provider{
		&credentials.EnvProvider{},
		&ec2rolecreds.EC2RoleProvider{
			Client: ec2metadata.New(sess),
		},
	}), nil
}
