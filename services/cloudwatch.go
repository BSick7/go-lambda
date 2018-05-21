package services

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func NewCloudwatch() (*cloudwatch.CloudWatch, error) {
	s, err := NewSession()
	if err != nil {
		return nil, err
	}
	return cloudwatch.New(s), nil
}
