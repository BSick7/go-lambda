package kinesis

import (
	"encoding/json"
	"reflect"
	"testing"
)

type fakeItem struct {
	partitionKey string
	Id           int    `json:"id"`
	Name         string `json:"name"`
}

func (f *fakeItem) PartitionKey() string {
	return f.partitionKey
}

func TestEventFromItems(t *testing.T) {
	pk1 := "1"
	pk2 := "2"

	got := EventFromItems([]*fakeItem{
		{
			partitionKey: pk1,
			Id:           1,
			Name:         "Test 1",
		},
		{
			partitionKey: pk2,
			Id:           2,
			Name:         "Test 2",
		},
	})

	want := &Event{
		Records: []*EventRecord{
			{
				EventSourceARN:    "arn:aws:kinesis:EXAMPLE",
				EventSource:       "aws:kinesis",
				AwsRegion:         "us-east-1",
				EventID:           "shardId-000000000000:49545115243490985018280067714973144582180062593244200961",
				EventVersion:      "1.0",
				InvokeIdentityArn: "arn:aws:iam::EXAMPLE",
				EventName:         "aws:kinesis:record",
				Kinesis: &EventRecordData{
					ApproximateArrivalTimestamp: 1428537600,
					KinesisSchemaVersion:        "1.0",
					SequenceNumber:              "49545115243490985018280067714973144582180062593244200961",
					PartitionKey:                &pk1,
					Data:                        "eyJpZCI6MSwibmFtZSI6IlRlc3QgMSJ9",
				},
			},
			{
				EventSourceARN:    "arn:aws:kinesis:EXAMPLE",
				EventSource:       "aws:kinesis",
				AwsRegion:         "us-east-1",
				EventID:           "shardId-000000000000:49545115243490985018280067714973144582180062593244200961",
				EventVersion:      "1.0",
				InvokeIdentityArn: "arn:aws:iam::EXAMPLE",
				EventName:         "aws:kinesis:record",
				Kinesis: &EventRecordData{
					ApproximateArrivalTimestamp: 1428537600,
					KinesisSchemaVersion:        "1.0",
					SequenceNumber:              "49545115243490985018280067714973144582180062593244200961",
					PartitionKey:                &pk2,
					Data:                        "eyJpZCI6MiwibmFtZSI6IlRlc3QgMiJ9",
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		gotRaw, _ := json.Marshal(got)
		wantRaw, _ := json.Marshal(want)
		t.Errorf("mismatched result\ngot  %s\nwant %s", string(gotRaw), string(wantRaw))
	}
}
