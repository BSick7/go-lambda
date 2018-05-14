package services

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoDB() (*dynamodb.DynamoDB, error) {
	s, err := NewSession()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(s), nil
}
