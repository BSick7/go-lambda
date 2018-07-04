package metric

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func CountPoint(name string, count int, dimensions ...Dimension) Point {
	return &countPoint{
		BasePoint: BasePoint{
			Name:       name,
			Unit:       cloudwatch.StandardUnitCount,
			Timestamp:  time.Now(),
			Dimensions: dimensions,
		},
		Count: count,
	}
}

type countPoint struct {
	BasePoint
	Count int
}

func (p *countPoint) ToAWS() cloudwatch.MetricDatum {
	md := p.BasePoint.ToAWS()
	md.Value = aws.Float64(float64(p.Count))
	return md
}
