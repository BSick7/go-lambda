package results

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	conf := aws.NewConfig()
	if region := os.Getenv("AWS_REGION"); region != "" {
		conf.WithRegion(region)
	}
	conf.WithCredentials(credentials.NewEnvCredentials())
	ses, err := session.NewSession(conf)
	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(ses)
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
