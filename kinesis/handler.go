package kinesis

import (
	"context"
	"encoding/json"

	"github.com/BSick7/go-lambda/scaffold"
)

type HandlerFunc func(ctx context.Context, evt *Event) (interface{}, error)

func Handler(fn HandlerFunc) scaffold.HandlerFunc {
	return func(ctx context.Context, payload []byte) ([]byte, error) {
		var evt Event
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		response, err := fn(ctx, &evt)
		if err != nil {
			return nil, err
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		return responseBytes, nil
	}
}
