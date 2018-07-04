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
		basePoint: basePoint{
			name:       name,
			unit:       cloudwatch.StandardUnitMicroseconds,
			dimensions: dimensions,
			timestamp:  time.Now(),
		},
		start: time.Now(),
	}
}

type durationPoint struct {
	basePoint
	start time.Time
	dur   *time.Duration
}

func (p *durationPoint) Mark() {
	if p.dur == nil {
		dur := time.Since(p.start).Round(time.Microsecond)
		p.dur = &dur
	}
}

func (p *durationPoint) ToAWS() cloudwatch.MetricDatum {
	p.Mark()
	md := p.basePoint.ToAWS()
	md.Value = aws.Float64(float64(p.dur.Nanoseconds() / 1000))
	return md
}
