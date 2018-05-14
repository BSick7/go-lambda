package dynamodb

import (
	"fmt"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      vals,
	})
	return err
}
