package cloudwatch

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type MetricPoint interface {
	ToAWS() *cloudwatch.MetricDatum
}

func CountMetricPoint(name string, count int) MetricPoint {
	return &countMetricPoint{
		name:  name,
		count: count,
		t:     time.Now(),
	}
}

type countMetricPoint struct {
	name  string
	count int
	t     time.Time
}

func (p *countMetricPoint) ToAWS() *cloudwatch.MetricDatum {
	return &cloudwatch.MetricDatum{
		MetricName:        aws.String(p.name),
		Value:             aws.Float64(float64(p.count)),
		StorageResolution: aws.Int64(1),
		Timestamp:         aws.Time(p.t),
	}
}
