package metric

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// NewDatadogEmitter creates an emitter that publishes custom metrics to datadog via API
func NewDatadogEmitter(apiKey, namespace string) Emitter {
	return &datadogEmitter{namespace: namespace}
}

type datadogEmitter struct {
	apiKey    string
	namespace string
	points    []*Point
}

func (e *datadogEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *datadogEmitter) Add(point *Point) {
	e.points = append(e.points, point)
}

func (e *datadogEmitter) Flush() error {
	body := datadogPost{Series: []datadogMetric{}}
	for _, point := range e.points {
		body.Series = append(body.Series, toDatadogMetric(point))
	}
	raw, _ := json.Marshal(body)
	_, err := http.Post(e.url(), "application/json", bytes.NewBuffer(raw))
	return err
}

func (e *datadogEmitter) url() string {
	return fmt.Sprintf("https://api.dat adoghq.com/api/v1/series?api_key=%s", e.apiKey)
}

type datadogPost struct {
	Series []datadogMetric `json:"series"`
}

type datadogMetric struct {
	Metric string      `json:"metric"`
	Unit   string      `json:"type"`
	Tags   []string    `json:"tags"`
	Points [][]float64 `json:"points"`
}

var datadogMetricTypes = map[string]string{
	"Count":        "count",
	"Microseconds": "microsecond",
}

func toDatadogMetric(point *Point) datadogMetric {
	tags := make([]string, 0)
	for n, v := range point.Tags {
		tags = append(tags, fmt.Sprintf("%s:%s", n, v))
	}

	return datadogMetric{
		Metric: point.Metric,
		Unit:   datadogMetricTypes[point.Unit],
		Tags:   tags,
		Points: [][]float64{
			{float64(point.Timestamp.Unix())},
		},
	}
}
