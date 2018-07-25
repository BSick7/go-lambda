package metric

import (
	"time"
)

type Point struct {
	Metric     string
	Unit       string
	Timestamp  time.Time
	Resolution *int64
	Value      float64
	Tags       map[string]string
}
