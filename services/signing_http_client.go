package services

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/sha1sum/aws_signing_client"
)

func NewAwsSigningHttpClient() (*http.Client, error) {
	creds, err := DefaultCredentials(DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("error creating credential chain: %s", err)
	}
	signer := v4.NewSigner(creds)
	awsClient, err := aws_signing_client.New(signer, nil, "es", "us-east-1")
	if err != nil {
		return nil, fmt.Errorf("error creating aws signing client: %s", err)
	}
	return awsClient, err
}
