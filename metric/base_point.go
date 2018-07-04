package metric

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type basePoint struct {
	name       string
	unit       cloudwatch.StandardUnit
	dimensions Dimensions
	resolution *int64
	timestamp  time.Time
}

func (p *basePoint) SetResolution(resolution int64) Point {
	p.resolution = &resolution
	return p
}

func (p *basePoint) ToAWS() cloudwatch.MetricDatum {
	return cloudwatch.MetricDatum{
		MetricName:        aws.String(p.name),
		Timestamp:         aws.Time(p.timestamp),
		Dimensions:        p.dimensions.ToAWS(),
		StorageResolution: p.resolution,
		Unit:              p.unit,
	}
}
