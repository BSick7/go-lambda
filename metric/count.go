package metric

import (
	"time"
)

func Count(name string, count int, tags map[string]string) *Point {
	return &Point{
		Metric:    name,
		Unit:      "Count",
		Timestamp: time.Now(),
		Value:     float64(count),
		Tags:      tags,
	}
}
