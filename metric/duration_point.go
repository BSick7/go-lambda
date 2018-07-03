package metric

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type DurationPointMarker interface {
	Point
	Mark()
}

// This will report duration in microseconds to cloudwatch
// The duration starts when this is created and ends when converted to MetricDatum via ToAWS
func DurationPoint(name string) DurationPointMarker {
	return &durationPoint{
		name:  name,
		start: time.Now(),
	}
}

type durationPoint struct {
	name  string
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
	return cloudwatch.MetricDatum{
		Unit:              cloudwatch.StandardUnitMicroseconds,
		MetricName:        aws.String(p.name),
		Value:             aws.Float64(float64(p.dur.Nanoseconds() / 1000)),
		StorageResolution: aws.Int64(1),
		Timestamp:         aws.Time(time.Now()),
	}
}
