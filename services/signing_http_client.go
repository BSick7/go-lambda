package services

import (
	"fmt"
	"net/http"

	"github.com/BSick7/aws_signing_client"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

func NewAwsSigningHttpClient() (*http.Client, error) {
	cfg, err := DefaultConfig()
	if err != nil {
		return nil, fmt.Errorf("error creating credential chain: %s", err)
	}
	signer := v4.NewSigner(cfg.Credentials)
	awsClient, err := aws_signing_client.New(signer, nil, "es", "us-east-1")
	if err != nil {
		return nil, fmt.Errorf("error creating aws signing client: %s", err)
	}
	return awsClient, err
}
