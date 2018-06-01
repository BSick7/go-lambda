package metric

import (
	"context"
	"log"
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
	log.Println(e.namespace, e.points)
	return nil
}
