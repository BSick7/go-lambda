package kinesis

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
)

func EventFromItems(items interface{}) *Event {
	t, s := reflect.TypeOf(items).Kind(), reflect.ValueOf(items)
	if t != reflect.Slice {
		panic("EventFromItems: items must be a slice")
	}

	evt := &Event{
		Records: []*EventRecord{},
	}
	for i := 0; i < s.Len(); i++ {
		d := s.Index(i).Interface()
		record := &EventRecord{
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
				PartitionKey:                getPartitionKey(d),
			},
		}
		record.Kinesis.JsonMarshal(d)
		evt.Records = append(evt.Records, record)
	}
	return evt
}

type Event struct {
	Records []*EventRecord `json:"Records"`
}

type EventRecord struct {
	EventSourceARN    string           `json:"eventSourceARN"`
	EventSource       string           `json:"eventSource"`
	AwsRegion         string           `json:"awsRegion"`
	EventID           string           `json:"eventID"`
	EventVersion      string           `json:"eventVersion"`
	Kinesis           *EventRecordData `json:"kinesis"`
	InvokeIdentityArn string           `json:"invokeIdentityARN"`
	EventName         string           `json:"eventName"`
}

type EventRecordData struct {
	ApproximateArrivalTimestamp float64 `json:"approximateArrivalTimestamp"`
	PartitionKey                *string `json:"partitionKey"`
	Data                        string  `json:"data"`
	KinesisSchemaVersion        string  `json:"kinesisSchemaVersion"`
	SequenceNumber              string  `json:"sequenceNumber"`
}

func (rk *EventRecordData) JsonMarshal(v interface{}) {
	raw, _ := json.Marshal(v)
	rk.Data = base64.StdEncoding.EncodeToString(raw)
}

func (rk EventRecordData) JsonUnmarshal(v interface{}) error {
	raw, _ := base64.StdEncoding.DecodeString(rk.Data)
	return json.Unmarshal(raw, v)
}
