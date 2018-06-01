package metric

import (
	"context"
)

func NewCloudwatchEmitter(namespace string) Emitter {
	return &cloudwatchEmitter{namespace: namespace}
}

type cloudwatchEmitter struct {
	namespace string
	points    []Point
}

func (e *cloudwatchEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *cloudwatchEmitter) Add(point Point) {
	e.points = append(e.points, point)
}

func (e *cloudwatchEmitter) Flush() error {
	return PutCloudwatch(e.namespace, e.points)
}
