package relay

import (
	"context"

	"github.com/BSick7/go-lambda/kinesis"
)

func NewKinesisEmitter(stream string) Emitter {
	return &kinesisEmitter{
		stream: stream,
	}
}

type kinesisEmitter struct {
	stream string
}

func (e *kinesisEmitter) Contextualize(ctx context.Context) context.Context {
	return context.WithValue(ctx, emitterContextKey{}, e)
}

func (e *kinesisEmitter) Emit(items interface{}) (int, error) {
	if e.stream == "" {
		return 0, nil
	}
	records, err := kinesis.Emit(e.stream, items)
	return len(records), err
}
