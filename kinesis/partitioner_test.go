package kinesis

import (
	"reflect"
	"testing"
)

func TestPartitioner_getPartitionKey(t *testing.T) {
	input1 := struct{}{}
	got := getPartitionKey(input1)
	var want interface{} = nil
	if reflect.DeepEqual(got, want) {
		t.Errorf("[0] mismatched partition key, got %+v, want %+v", got, want)
	}

	input2 := struct{ PartitionKey func() string }{PartitionKey: func() string { return "key-1" }}
	got = getPartitionKey(input2)
	want = "key-1"
	if reflect.DeepEqual(got, want) {
		t.Errorf("[1] mismatched partition key, got %+v, want %+v", got, want)
	}
}
