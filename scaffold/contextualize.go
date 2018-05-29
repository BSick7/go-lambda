package scaffold

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Contextualizer interface {
	Contextualize(ctx context.Context) context.Context
}

type HandlerFunc func(ctx context.Context, payload []byte) ([]byte, error)

func Contextualize(handler HandlerFunc, contextualizers ...Contextualizer) lambda.Handler {
	return contextualHandler{handler, contextualizers}
}

type contextualHandler struct {
	fn              HandlerFunc
	contextualizers []Contextualizer
}

func (handler contextualHandler) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	for _, ctxlizer := range handler.contextualizers {
		ctx = ctxlizer.Contextualize(ctx)
	}
	return handler.fn(ctx, payload)
}
