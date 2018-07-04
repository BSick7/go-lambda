package metric

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func CountPoint(name string, count int, dimensions ...Dimension) Point {
	return &countPoint{
		basePoint: basePoint{
			name:       name,
			unit:       cloudwatch.StandardUnitCount,
			timestamp:  time.Now(),
			dimensions: dimensions,
		},
		count: count,
	}
}

type countPoint struct {
	basePoint
	count int
}

func (p *countPoint) ToAWS() cloudwatch.MetricDatum {
	md := p.basePoint.ToAWS()
	md.Value = aws.Float64(float64(p.count))
	return md
}
