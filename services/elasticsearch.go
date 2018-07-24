package services

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/BSick7/aws-signing/signing"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/olivere/elastic"
)

type ElasticsearchConfig struct {
	EndpointUrl          string
	UseAwsRequestSigning bool
	RequestTimeout       time.Duration
}

func (c ElasticsearchConfig) Scheme() string {
	u, _ := url.Parse(c.EndpointUrl)
	scheme := u.Scheme
	if scheme == "" {
		scheme = "https"
	}
	return scheme
}

func (c ElasticsearchConfig) ClientOptions() ([]elastic.ClientOptionFunc, error) {
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(c.EndpointUrl),
		elastic.SetScheme(c.Scheme()),
	}
	if c.UseAwsRequestSigning {
		cfg, err := DefaultConfig()
		if err != nil {
			return nil, fmt.Errorf("error creating aws configuration: %s", err)
		}
		signer := v4.NewSigner(cfg.Credentials)
		httpClient := &http.Client{Transport: signing.NewTransport(signer, "es", cfg.Region)}
		opts = append(opts, elastic.SetHttpClient(httpClient))
	}
	return opts, nil
}
