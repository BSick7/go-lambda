package metric

import (
	"context"
	"fmt"
	"strings"
)

type MultiEmitter []Emitter

func (e MultiEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e MultiEmitter) Add(point Point) {
	for _, item := range e {
		item.Add(point)
	}
}

func (e MultiEmitter) Flush() error {
	errs := make([]string, 0)
	for _, item := range e {
		if err := item.Flush(); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("multiple metric errors:\n  %s", strings.Join(errs, "\n  "))
	}
	return nil
}
