package metric

import (
	"context"
	"log"

	"github.com/BSick7/go-lambda/scaffold"
)

type Emitter interface {
	scaffold.Contextualizer
	Add(point *Point)
	Flush() error
}

type emitterContextKey struct{}

func ContextEmitter(ctx context.Context) Emitter {
	emitter, _ := ctx.Value(emitterContextKey{}).(Emitter)
	return emitter
}

func WithEmitter(ctx context.Context, e Emitter) context.Context {
	return context.WithValue(ctx, emitterContextKey{}, e)
}

func FlushAndLog(emitter Emitter) {
	if err := emitter.Flush(); err != nil {
		log.Println(err)
	}
}
