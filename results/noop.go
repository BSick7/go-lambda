package results

import (
	"context"
)

func NewNoopSaver() Saver {
	return &noopSaver{}
}

type noopSaver struct{}

func (s *noopSaver) Contextualize(ctx context.Context) context.Context { return WithSaver(ctx, s) }
func (s *noopSaver) Save(key string, data []byte) (string, error)      { return "", nil }
