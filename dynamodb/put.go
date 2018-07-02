package dynamodb

import (
	"fmt"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

func Put(tableName string, item interface{}) error {
	vals, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	svc, err := services.NewDynamoDB()
	if err != nil {
		return fmt.Errorf("could not create dynamo db service: %s", err)
	}

	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      vals,
	})
	_, err = req.Send()
	return err
}
