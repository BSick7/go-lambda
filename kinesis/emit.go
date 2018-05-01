package kinesis

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// Emit a slice of data to a kinesis stream
// This will convert the data into record entries and emit properly
// By default, the partition key is not set
// If an item implements Partitioner, partition key will be retrieved accordingly
func Emit(stream string, items interface{}) ([]*kinesis.PutRecordsResultEntry, error) {
	svc, err := services.NewKinesis()
	if err != nil {
		return nil, fmt.Errorf("could not create kinesis service: %s", err)
	}

	records, err := buildRecords(items)
	if err != nil {
		return nil, fmt.Errorf("could not create kinesis records: %s", err)
	}

	out, err := svc.PutRecords(&kinesis.PutRecordsInput{
		StreamName: aws.String(stream),
		Records:    records,
	})
	if err != nil {
		return nil, fmt.Errorf("could not put kinesis records to %q: %s", stream, err)
	}
	return out.Records, nil
}

func buildRecords(items interface{}) ([]*kinesis.PutRecordsRequestEntry, error) {
	t, s := reflect.TypeOf(items).Kind(), reflect.ValueOf(items)
	switch t {
	default:
		return nil, fmt.Errorf("expected Emit items input to be slice, got %s", t)
	case reflect.Slice:
	}

	records := make([]*kinesis.PutRecordsRequestEntry, s.Len())
	for i := 0; i < s.Len(); i++ {
		item := s.Index(i).Interface()
		partitionKey := getPartitionKey(item)
		raw, _ := json.Marshal(item)
		records[i] = &kinesis.PutRecordsRequestEntry{
			Data:         raw,
			PartitionKey: partitionKey,
		}
	}
	return records, nil
}
