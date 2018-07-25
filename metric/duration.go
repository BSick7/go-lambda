package metric

import (
	"time"
)

// AddFuncDuration enables easy reporting of function duration
// In the following example, "execution-duration" is added to metrics
//   and "marked" at the end of the function to record function duration
//
// defer AddFuncDuration(metrics, "execution-duration", nil)
func AddFuncDuration(emitter Emitter, name string, tags map[string]string) func() {
	p := Duration(name, tags)
	emitter.Add(&p.Point)
	return p.Mark
}

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
