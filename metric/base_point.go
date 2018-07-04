package metric

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type BasePoint struct {
	Name       string
	Unit       cloudwatch.StandardUnit
	Dimensions Dimensions
	Resolution *int64
	Timestamp  time.Time
}

func (p *BasePoint) SetResolution(resolution int64) Point {
	p.Resolution = &resolution
	return p
}

func (p *BasePoint) ToAWS() cloudwatch.MetricDatum {
	return cloudwatch.MetricDatum{
		MetricName:        aws.String(p.Name),
		Timestamp:         aws.Time(p.Timestamp),
		Dimensions:        p.Dimensions.ToAWS(),
		StorageResolution: p.Resolution,
		Unit:              p.Unit,
	}
}
