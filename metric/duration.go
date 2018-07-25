package metric

import (
	"time"
)

// This will report duration in microseconds to cloudwatch
// The duration starts when this is created and ends when converted to MetricDatum via ToAWS
func Duration(name string, tags map[string]string) *DurationPoint {
	return &DurationPoint{
		Point: Point{
			Metric:    name,
			Unit:      "Microseconds",
			Timestamp: time.Now(),
			Tags:      tags,
		},
		start: time.Now(),
	}
}

type DurationPoint struct {
	Point
	start time.Time
}

func (p *DurationPoint) Mark() {
	if p.Value <= 0 {
		dur := time.Since(p.start).Round(time.Microsecond)
		p.Value = float64(dur.Nanoseconds() / 1000)
	}
}
