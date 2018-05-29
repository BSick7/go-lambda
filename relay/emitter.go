package relay

import (
	"context"

	"github.com/BSick7/go-lambda/scaffold"
)

type emitterContextKey struct{}

func ContextEmitter(ctx context.Context) Emitter {
	emitter, _ := ctx.Value(emitterContextKey{}).(Emitter)
	return emitter
}

type Emitter interface {
	scaffold.Contextualizer
	Emit(items interface{}) (int, error)
}
