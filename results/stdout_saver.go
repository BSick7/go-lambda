package results

import (
	"context"
	"fmt"
)

func NewStdoutSaver() Saver {
	return &stdoutSaver{}
}

type stdoutSaver struct {
}

func (s *stdoutSaver) Contextualize(ctx context.Context) context.Context {
	return context.WithValue(ctx, saverContextKey{}, s)
}

func (s *stdoutSaver) Save(key string, data []byte) (string, error) {
	fmt.Printf("%s:\n%s\n", key, string(data))
	return "", nil
}
