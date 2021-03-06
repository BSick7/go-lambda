package dynamodb

import (
	"fmt"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Get(tableName string, keyName string, key string) (map[string]dynamodb.AttributeValue, error) {
	svc, err := services.NewDynamoDB()
	if err != nil {
		return nil, fmt.Errorf("could not create dynamodb service: %s", err)
	}

	req := svc.QueryRequest(&dynamodb.QueryInput{
		Limit:                  aws.Int64(1),
		ConsistentRead:         aws.Bool(true),
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("#%s = :value", keyName)),
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":value": {
				S: aws.String(key),
			},
		},
		ExpressionAttributeNames: map[string]string{
			"#" + keyName: keyName,
		},
	})
	out, err := req.Send()
	if err != nil {
		return nil, err
	}

	if out == nil || *out.Count <= 0 {
		return nil, nil
	}

	return out.Items[0], nil
}
