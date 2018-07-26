package metric

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestDatadogEmitter_Flush(t *testing.T) {
	var got datadogPost
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/v1/series" {
			defer r.Body.Close()
			if raw, err := ioutil.ReadAll(r.Body); err != nil {
				t.Errorf("unable to read body: %s", err)
			} else if err := json.Unmarshal(raw, &got); err != nil {
				t.Errorf("unable to unmarshal body: %s", err)
			}
		} else {
			http.Error(w, "unknown endpoint", http.StatusInternalServerError)
		}
	}))
	defer server.Close()

	e := datadogEmitter{
		server:    server.URL,
		apiKey:    "xyz",
		namespace: "ns",
		allPoints: map[string][]*Point{},
	}
	e.Add(stripPointTimestamp(Count("some-count", 1, nil)))
	someDuration := Duration("some-duration", nil)
	someDuration.Value = 1500
	e.Add(stripPointTimestamp(&someDuration.Point))
	e.Add(stripPointTimestamp(Count("some-count", 5, nil)))

	if err := e.Flush(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	want := datadogPost{
		Series: []datadogMetric{
			{
				Metric: "ns.some.count",
				Unit:   "count",
				Tags:   []string{},
				Points: [][]float64{
					{0, 1},
					{0, 5},
				},
			},
			{
				Metric: "ns.some.duration",
				Unit:   "gauge",
				Tags:   []string{},
				Points: [][]float64{
					{0, 1500},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("mismatched expectations\ngot\n%+v\nwant\n%+v", got, want)
	}
}

func stripPointTimestamp(point *Point) *Point {
	return &Point{
		Metric:    point.Metric,
		Unit:      point.Unit,
		Tags:      point.Tags,
		Timestamp: time.Unix(0, 0),
		Value:     point.Value,
	}
}
