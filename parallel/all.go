package parallel

import (
	"sync"
)

type AllEngine interface {
	Add(runner Runner)
	Await() ([]interface{}, []error)
}

type allEngine struct {
	wg      *sync.WaitGroup
	mu      *sync.Mutex
	results []interface{}
	errs    []error
}

func All() AllEngine {
	return &allEngine{
		wg:      &sync.WaitGroup{},
		mu:      &sync.Mutex{},
		results: []interface{}{},
		errs:    []error{},
	}
}

func (e *allEngine) Add(runner Runner) {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		result, err := runner.Run()

		e.mu.Lock()
		defer e.mu.Unlock()
		if err != nil {
			e.errs = append(e.errs, err)
		} else {
			e.results = append(e.results, result)
		}
	}()
}

func (e *allEngine) Await() ([]interface{}, []error) {
	e.wg.Wait()
	return e.results, e.errs
}
