package metric

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

// This will report duration in microseconds to cloudwatch
// The duration starts when this is created and ends when converted to MetricDatum via ToAWS
func DurationPoint(name string, dimensions ...Dimension) PointMarker {
	return &durationPoint{
		BasePoint: BasePoint{
			Name:       name,
			Unit:       cloudwatch.StandardUnitMicroseconds,
			Dimensions: dimensions,
			Timestamp:  time.Now(),
		},
		start: time.Now(),
	}
}

type durationPoint struct {
	BasePoint
	start    time.Time
	Duration *time.Duration
}

func (p *durationPoint) Mark() {
	if p.Duration == nil {
		dur := time.Since(p.start).Round(time.Microsecond)
		p.Duration = &dur
	}
}

func (p *durationPoint) ToAWS() cloudwatch.MetricDatum {
	p.Mark()
	md := p.BasePoint.ToAWS()
	md.Value = aws.Float64(float64(p.Duration.Nanoseconds() / 1000))
	return md
}
