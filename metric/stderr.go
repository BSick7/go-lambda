package metric

import (
	"context"
	"encoding/json"
	"log"
)

// NewStderrEmitter creates an emitter dumps metrics to stderr
func NewStderrEmitter(namespace string) Emitter {
	return &stderrEmitter{namespace: namespace}
}

type stderrEmitter struct {
	namespace string
	points    []*Point
}

func (e *stderrEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *stderrEmitter) Add(point *Point) {
	e.points = append(e.points, point)
}

func (e *stderrEmitter) Flush() error {
	raw, _ := json.Marshal(e.points)
	log.Println(e.namespace, string(raw))
	return nil
}
