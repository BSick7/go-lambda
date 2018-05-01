package services

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession() (*session.Session, error) {
	conf := aws.NewConfig()
	if region := os.Getenv("AWS_REGION"); region != "" {
		conf.WithRegion(region)
	}
	conf.WithCredentials(credentials.NewEnvCredentials())
	return session.NewSession(conf)
}
