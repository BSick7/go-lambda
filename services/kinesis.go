package services

import "github.com/aws/aws-sdk-go/service/kinesis"

func NewKinesis() (*kinesis.Kinesis, error) {
	s, err := NewSession()
	if err != nil {
		return nil, err
	}
	return kinesis.New(s), nil
}
