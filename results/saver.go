package results

import (
	"context"

	"github.com/BSick7/go-lambda/scaffold"
)

type Saver interface {
	scaffold.Contextualizer
	Save(key string, data []byte) (string, error)
}

type saverContextKey struct{}

func ContextSaver(ctx context.Context) Saver {
	emitter, _ := ctx.Value(saverContextKey{}).(Saver)
	return emitter
}
