package services

import (
	"fmt"
	"net/url"
	"time"

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
		if httpClient, err := NewAwsSigningHttpClient(); err != nil {
			return nil, fmt.Errorf("error creating elasticsearch http client: %s", err)
		} else {
			opts = append(opts, elastic.SetHttpClient(httpClient))
		}
	}
	return opts, nil
}
