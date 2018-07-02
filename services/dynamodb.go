package services

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDB() (*dynamodb.DynamoDB, error) {
	cfg, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(cfg), nil
}
