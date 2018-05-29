package scaffold

import (
	"context"
)

func BuildContext(contextualizers ...Contextualizer) context.Context {
	ctx := context.Background()
	for _, ctxlizer := range contextualizers {
		ctx = ctxlizer.Contextualize(ctx)
	}
	return ctx
}
