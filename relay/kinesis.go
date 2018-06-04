package relay

import (
	"context"

	"github.com/BSick7/go-lambda/kinesis"
)

func NewKinesisEmitter(stream string) Emitter {
	return &kinesisEmitter{
		stream: stream,
		items:  make([]interface{}, 0),
	}
}

type kinesisEmitter struct {
	stream string
	items  []interface{}
}

func (e *kinesisEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *kinesisEmitter) Emit(item interface{}) error {
	e.items = append(e.items, item)
	if len(e.items) >= 500 {
		return e.Flush()
	}
	return nil
}

func (e *kinesisEmitter) Flush() error {
	if len(e.items) <= 0 {
		return nil
	}
	_, err := kinesis.Emit(e.stream, e.items)
	e.items = make([]interface{}, 0)
	return err
}
