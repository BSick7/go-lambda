package main

import (
	"fmt"

	"github.com/BSick7/go-lambda/kinesis"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(event *kinesis.Event) error {
	for _, record := range event.Records {
		fmt.Printf("partition key: %+v\n", record.Kinesis.PartitionKey)
		data := map[string]interface{}{}
		if err := record.Kinesis.JsonUnmarshal(&data); err != nil {
			fmt.Printf("error reading record: %s", err)
		} else {
			fmt.Printf("record: %+v", data)
		}
	}
	return nil
}
