package cloudwatch

import (
	"log"
)

type MetricEmitter interface {
	Add(point MetricPoint)
	Flush() error
}

func NewMetricEmitter(namespace string) MetricEmitter {
	return &metricEmitter{namespace: namespace}
}

type metricEmitter struct {
	namespace string
	points    []MetricPoint
}

func (e *metricEmitter) Add(point MetricPoint) {
	e.points = append(e.points, point)
}

func (e *metricEmitter) Flush() error {
	return PutMetrics(e.namespace, e.points)
}

func NewStderrEmitter(namespace string) MetricEmitter {
	return &stderrEmitter{namespace: namespace}
}

type stderrEmitter struct {
	namespace string
	points    []MetricPoint
}

func (e *stderrEmitter) Add(point MetricPoint) {
	e.points = append(e.points, point)
}

func (e *stderrEmitter) Flush() error {
	log.Println(e.namespace, e.points)
	return nil
}
