package parallel

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type interfaceSlice []interface{}

func (s interfaceSlice) Len() int { return len(s) }
func (s interfaceSlice) Less(i, j int) bool {
	return fmt.Sprintf("%+v", s[i]) < fmt.Sprintf("%+v", s[j])
}
func (s interfaceSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type errorSlice []error

func (s errorSlice) Len() int           { return len(s) }
func (s errorSlice) Less(i, j int) bool { return s[i].Error() < s[j].Error() }
func (s errorSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func TestAll(t *testing.T) {
	runners := []Runner{
		mockRunner{1, nil},
		mockRunner{2, nil},
		mockRunner{nil, fmt.Errorf("invalid id")},
		mockRunner{nil, fmt.Errorf("invalid name")},
	}

	engine := All()
	for _, runner := range runners {
		engine.Add(runner)
	}
	gotResults, gotErrs := engine.Await()
	wantResults := []interface{}{1, 2}
	wantErrs := []error{errors.New("invalid id"), errors.New("invalid name")}

	sort.Sort(interfaceSlice(gotResults))
	sort.Sort(errorSlice(gotErrs))

	if !reflect.DeepEqual(gotResults, wantResults) {
		t.Errorf("mismatched result\ngot  %+v\nwant %+v", gotResults, wantResults)
	}

	if !reflect.DeepEqual(gotErrs, wantErrs) {
		t.Errorf("mismatched result\ngot  %s\nwant %s", gotErrs, wantErrs)
	}
}
