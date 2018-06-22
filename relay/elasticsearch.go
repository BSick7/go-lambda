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
		items:                map[string][]elastic.BulkableRequest{},
	}
}

type elasticsearchEmitter struct {
	endpointUrl          string
	useAwsRequestSigning bool
	indexer              ElasticsearchIndexer
	client               *elastic.Client
	items                map[string][]elastic.BulkableRequest
	sync.Mutex
}

func (e *elasticsearchEmitter) Contextualize(ctx context.Context) context.Context {
	return WithEmitter(ctx, e)
}

func (e *elasticsearchEmitter) Emit(item interface{}) error {
	index := e.indexer.Index(item)
	req := elastic.NewBulkIndexRequest().
		Index(index).
		Doc(e.indexer.BodyJson(item))

	if indextype := e.indexer.Type(item); indextype != nil {
		req = req.Type(*indextype)
	}
	if id := e.indexer.Id(item); id != nil {
		req = req.Id(*id)
	}

	e.Lock()
	defer e.Unlock()
	if grp, ok := e.items[index]; !ok {
		e.items[index] = []elastic.BulkableRequest{req}
	} else {
		e.items[index] = append(grp, req)
	}

	return nil
}

func (e *elasticsearchEmitter) Flush() error {
	if err := e.init(); err != nil {
		return err
	}

	bulks := make([]*elastic.BulkService, 0)
	func() {
		e.Lock()
		defer e.Unlock()
		for index, grp := range e.items {
			bulk := e.client.Bulk().Index(index).
				Add(grp...)
			bulks = append(bulks, bulk)
		}
		e.items = map[string][]elastic.BulkableRequest{}
	}()

	for _, bulk := range bulks {
		if _, err := bulk.Do(context.Background()); err != nil {
			return err
		}
	}
	return nil
}

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
