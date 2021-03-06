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

func NewElasticsearchEmitter(cfg services.ElasticsearchConfig, indexer ElasticsearchIndexer) Emitter {
	return &elasticsearchEmitter{
		cfg:     cfg,
		indexer: indexer,
		items:   map[string][]elastic.BulkableRequest{},
	}
}

type elasticsearchEmitter struct {
	cfg     services.ElasticsearchConfig
	indexer ElasticsearchIndexer
	client  *elastic.Client
	items   map[string][]elastic.BulkableRequest
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

	ctx := context.Background()
	if e.cfg.RequestTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), e.cfg.RequestTimeout)
		defer cancel()
	}

	for _, bulk := range bulks {
		if _, err := bulk.Do(ctx); err != nil {
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

	opts, err := e.cfg.ClientOptions()
	if err != nil {
		return err
	}
	opts = append(opts, elastic.SetSniff(false))

	client, err := elastic.NewClient(opts...)
	if err != nil {
		return fmt.Errorf("error creating elasticsearch client: %s", err)
	}
	e.client = client
	return nil
}
