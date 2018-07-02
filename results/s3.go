package results

import (
	"bytes"
	"context"
	"fmt"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

func NewS3Saver(bucket string) Saver {
	return &s3Saver{
		bucket: bucket,
	}
}

type s3Saver struct {
	bucket string
}

func (s *s3Saver) Contextualize(ctx context.Context) context.Context {
	return WithSaver(ctx, s)
}

func (s *s3Saver) Save(key string, data []byte) (string, error) {
	cfg, err := services.DefaultConfig()
	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(cfg)
	out, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ContentType: aws.String("text/plain"),
		Body:        bytes.NewBuffer(data),
	})
	if err != nil {
		return "", fmt.Errorf("error saving %q: %s", key, err)
	}
	return out.Location, nil

}

func (s *s3Saver) AlwaysSave() bool {
	return false
}
