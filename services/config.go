package services

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func DefaultConfig() (aws.Config, error) {
	conf, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return aws.Config{}, err
	}
	if region := os.Getenv("AWS_REGION"); region != "" {
		conf.Region = region
	}
	return conf, nil
}
