package services

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
)

func DefaultConfig() *aws.Config {
	conf := aws.NewConfig()
	if region := os.Getenv("AWS_REGION"); region != "" {
		conf.WithRegion(region)
	}
	return conf
}
