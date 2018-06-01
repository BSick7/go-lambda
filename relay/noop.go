package relay

import (
	"context"
)

func NewNoopEmitter() Emitter {
	return &noopEmitter{}
}

type noopEmitter struct {
}

func (e *noopEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *noopEmitter) Emit(items interface{}) (int, error) {
	return 0, nil
}
