package kinesis

type Partitioner interface {
	PartitionKey() string
}

func getPartitionKey(item interface{}) *string {
	if p, ok := item.(Partitioner); ok {
		pk := p.PartitionKey()
		return &pk
	}
	return nil
}
