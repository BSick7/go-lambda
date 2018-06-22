package relay

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/BSick7/go-lambda/services"
	"github.com/olivere/elastic"
	"net/http"
)

type ElasticsearchIndexer struct {
	Index    func(item interface{}) string
	Type     func(item interface{}) *string
	Id       func(item interface{}) *string
	BodyJson func(item interface{}) interface{}
}

func NewElasticsearchEmitter(endpointUrl string, useAwsRequestSigning bool, indexer ElasticsearchIndexer) Emitter {
	return &elasticsearchEmitter{
		endpointUrl:          endpointUrl,
		useAwsRequestSigning: useAwsRequestSigning,
		indexer:              indexer,
	}
}

type elasticsearchEmitter struct {
	endpointUrl          string
	useAwsRequestSigning bool
	indexer              ElasticsearchIndexer
	client               *elastic.Client
	sync.Mutex
}

func (e *elasticsearchEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *elasticsearchEmitter) Emit(item interface{}) error {
	if err := e.init(); err != nil {
		return err
	}

	svc := e.client.Index().
		Index(e.indexer.Index(item)).
		BodyJson(e.indexer.BodyJson(item))

	if indextype := e.indexer.Type(item); indextype != nil {
		svc = svc.Type(*indextype)
	}
	if id := e.indexer.Id(item); id != nil {
		svc = svc.Id(*id)
	}

	_, err := svc.Do(context.Background())
	if err != nil {
		return fmt.Errorf("error emitting index: %s", err)
	}
	return nil
}

func (e *elasticsearchEmitter) Flush() error { return nil }

func (e *elasticsearchEmitter) init() error {
	e.Lock()
	defer e.Unlock()

	if e.client != nil {
		return nil
	}

	httpClient, err := e.getHttpClient()
	if err != nil {
		return fmt.Errorf("error creating elasticsearch http client: %s", err)
	}

	scheme := ""
	if u, err := url.Parse(e.endpointUrl); err == nil {
		scheme = u.Scheme
	}
	if scheme == "" {
		scheme = "https"
	}

	client, err := elastic.NewClient(
		elastic.SetURL(e.endpointUrl),
		elastic.SetScheme(scheme),
		elastic.SetSniff(false),
		elastic.SetHttpClient(httpClient),
	)
	if err != nil {
		return fmt.Errorf("error creating elasticsearch client: %s", err)
	}
	e.client = client
	return nil
}

func (e *elasticsearchEmitter) getHttpClient() (*http.Client, error) {
	if e.useAwsRequestSigning {
		return services.NewAwsSigningHttpClient()
	}
	return http.DefaultClient, nil
}
