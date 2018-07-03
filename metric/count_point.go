package metric

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func CountPoint(name string, count int) Point {
	return &countPoint{
		name:  name,
		count: count,
		t:     time.Now(),
	}
}

type countPoint struct {
	name  string
	count int
	t     time.Time
}

func (p *countPoint) ToAWS() cloudwatch.MetricDatum {
	return cloudwatch.MetricDatum{
		Unit:              cloudwatch.StandardUnitCount,
		MetricName:        aws.String(p.name),
		Value:             aws.Float64(float64(p.count)),
		StorageResolution: aws.Int64(1),
		Timestamp:         aws.Time(p.t),
	}
}
