package metric

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// NewDatadogEmitter creates an emitter that publishes custom metrics to datadog via API
func NewDatadogEmitter(apiKey, namespace string) Emitter {
	return &datadogEmitter{
		server:    "https://api.datadoghq.com/api",
		apiKey:    apiKey,
		namespace: namespace,
	}
}

type datadogEmitter struct {
	server    string
	apiKey    string
	namespace string
	allPoints map[string][]*Point
}

func (e *datadogEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *datadogEmitter) Add(point *Point) {
	name := e.metricName(point)
	if pts, ok := e.allPoints[name]; ok {
		e.allPoints[name] = append(pts, point)
	} else {
		e.allPoints[name] = []*Point{point}
	}
}

func (e *datadogEmitter) Flush() error {
	body := datadogPost{Series: []datadogMetric{}}
	allPoints := e.allPoints
	e.allPoints = map[string][]*Point{}
	for name, points := range allPoints {
		body.Series = append(body.Series, toDatadogMetric(name, points))
	}

	raw, _ := json.Marshal(body)
	_, err := http.Post(e.url(), "application/json", bytes.NewBuffer(raw))
	return err
}

func (e *datadogEmitter) metricName(point *Point) string {
	metric := point.Metric
	if e.namespace != "" {
		metric = fmt.Sprintf("%s.%s", e.namespace, metric)
	}
	return strings.Replace(strings.Replace(metric,
		"-", ".", -1),
		"/", ".", -1)
}

func (e *datadogEmitter) url() string {
	return fmt.Sprintf("%s/v1/series?api_key=%s", e.server, e.apiKey)
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
	"Microseconds": "gauge",
}

func toDatadogMetric(name string, points []*Point) datadogMetric {
	m := datadogMetric{
		Metric: name,
		Unit:   "count",
		Tags:   []string{},
		Points: [][]float64{},
	}
	if len(points) < 1 {
		return m
	}

	first := points[0]
	m.Unit = datadogMetricTypes[first.Unit]
	for n, v := range first.Tags {
		m.Tags = append(m.Tags, fmt.Sprintf("%s:%s", n, v))
	}

	for _, pt := range points {
		m.Points = append(m.Points, []float64{
			float64(pt.Timestamp.Unix()),
			pt.Value,
		})
	}

	return m
}
