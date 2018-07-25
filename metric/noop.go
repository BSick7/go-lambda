package metric

import (
	"context"
)

// NewNoopEmitter creates an emitter that ignores metrics
func NewNoopEmitter() Emitter { return &noopEmitter{} }

type noopEmitter struct{}

func (e *noopEmitter) Contextualize(ctx context.Context) context.Context { return WithEmitter(ctx, e) }
func (e *noopEmitter) Add(point *Point)                                  {}
func (e *noopEmitter) Flush() error                                      { return nil }
