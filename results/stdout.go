package results

import (
	"context"
	"fmt"
)

func NewStdoutSaver(alwaysSave bool) Saver {
	return &stdoutSaver{alwaysSave: alwaysSave}
}

type stdoutSaver struct {
	alwaysSave bool
}

func (s *stdoutSaver) Contextualize(ctx context.Context) context.Context {
	return WithSaver(ctx, s)
}

func (s *stdoutSaver) Save(key string, data []byte) (string, error) {
	fmt.Printf("%s:\n%s\n", key, string(data))
	return "", nil
}

func (s *stdoutSaver) AlwaysSave() bool {
	return s.alwaysSave
}
