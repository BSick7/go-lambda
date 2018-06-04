package relay

import (
	"context"
)

func NewNoopEmitter() Emitter {
	return &noopEmitter{}
}

type noopEmitter struct{}

func (e *noopEmitter) Contextualize(ctx context.Context) context.Context { return WithEmitter(ctx, e) }
func (e *noopEmitter) Emit(item interface{}) error                       { return nil }
func (e *noopEmitter) Flush() error                                      { return nil }
