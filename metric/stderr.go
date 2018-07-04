package metric

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func NewStderrEmitter(namespace string) Emitter {
	return &stderrEmitter{namespace: namespace}
}

type stderrEmitter struct {
	namespace string
	points    []Point
}

func (e *stderrEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *stderrEmitter) Add(point Point) {
	e.points = append(e.points, point)
}

func (e *stderrEmitter) Flush() error {
	pts := make([]cloudwatch.MetricDatum, 0)
	for _, point := range e.points {
		pts = append(pts, point.ToAWS())
	}
	raw, _ := json.Marshal(pts)
	log.Println(e.namespace, string(raw))
	return nil
}
