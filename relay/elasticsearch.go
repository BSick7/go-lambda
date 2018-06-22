package relay

import (
	"context"
	"fmt"
	"sync"

	"github.com/BSick7/go-lambda/services"
	"github.com/olivere/elastic"
)

type ElasticsearchIndexer struct {
	Index    func(item interface{}) string
	Type     func(item interface{}) *string
	Id       func(item interface{}) *string
	BodyJson func(item interface{}) interface{}
}

func NewElasticsearchEmitter(endpoint string, indexer ElasticsearchIndexer) Emitter {
	return &elasticsearchEmitter{
		endpoint: endpoint,
		indexer:  indexer,
	}
}

type elasticsearchEmitter struct {
	endpoint string
	indexer  ElasticsearchIndexer
	client   *elastic.Client
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

	httpClient, err := services.NewAwsSigningHttpClient()
	if err != nil {
		return err
	}

	client, err := elastic.NewClient(
		elastic.SetURL(e.endpoint),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
		elastic.SetHttpClient(httpClient),
	)
	if err != nil {
		return fmt.Errorf("error creating elasticsearch client: %s", err)
	}
	e.client = client
	return nil
}
