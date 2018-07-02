package services

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func NewCloudwatch() (*cloudwatch.CloudWatch, error) {
	cfg, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return cloudwatch.New(cfg), nil
}
