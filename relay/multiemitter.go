package relay

import (
	"context"
	"errors"
	"strings"
)

type multiEmitter struct {
	Emitters []Emitter
}

func NewMultiEmitter(emitters ...Emitter) Emitter {
	return &multiEmitter{
		Emitters: emitters,
	}
}

func (e *multiEmitter) Contextualize(ctx context.Context) context.Context { return WithEmitter(ctx, e) }

func (e *multiEmitter) Emit(item interface{}) error {
	var errs []string
	for _, emitter := range e.Emitters {
		if err := emitter.Emit(item); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New("emit errors:\n    " + strings.Join(errs, "\n    "))
	}
	return nil
}

func (e *multiEmitter) Flush() error {
	var errs []string
	for _, emitter := range e.Emitters {
		if err := emitter.Flush(); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New("flush errors:\n    " + strings.Join(errs, "\n    "))
	}
	return nil
}
