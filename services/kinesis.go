package services

import "github.com/aws/aws-sdk-go-v2/service/kinesis"

func NewKinesis() (*kinesis.Kinesis, error) {
	cfg, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return kinesis.New(cfg), nil
}
