package prometheus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/parse"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultScheme = "http"
	defaultPath   = "/metrics"
)

var (
	// HostParser validates Prometheus URLs
	HostParser = parse.URLHostParserBuilder{
		DefaultScheme: defaultScheme,
		DefaultPath:   defaultPath,
	}.Build()
)

func MetricSetBuilder(mapping *MetricsMapping) func(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return func(base mb.BaseMetricSet) (mb.MetricSet, error) {
		prometheus, err := NewPrometheusClient(base)
		if err != nil {
			return nil, err
		}
		return &prometheusMetricSet{
			BaseMetricSet: base,
			prometheus:    prometheus,
			mapping:       mapping,
		}, nil
	}
}

type prometheusMetricSet struct {
	mb.BaseMetricSet
	prometheus Prometheus
	mapping    *MetricsMapping
}

func (m *prometheusMetricSet) Fetch() ([]common.MapStr, error) {
	return m.prometheus.GetProcessedMetrics(m.mapping)
}

// TestCases holds the list of test cases to test a metricset
type TestCases []struct {
	// MetricsFile containing Prometheus outputted metrics
	MetricsFile string

	// ExpectedFile containing resulting documents
	ExpectedFile string
}

// TestMetricSet goes over the given TestCases and ensures that source Prometheus metrics gets converted into the expected
// events when passed by the given metricset.
// If update is true, the expected JSON file will be updated with the result
func TestMetricSet(t *testing.T, module, metricset string, cases TestCases, update bool) {
	for _, test := range cases {
		t.Logf("Testing %s file\n", test.MetricsFile)

		file, err := os.Open(test.MetricsFile)
		assert.NoError(t, err, "cannot open test file "+test.MetricsFile)

		body, err := ioutil.ReadAll(file)
		assert.NoError(t, err, "cannot read test file "+test.MetricsFile)

		server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/plain; charset=ISO-8859-1")
			w.Write([]byte(body))
		}))

		server.Start()
		defer server.Close()

		config := map[string]interface{}{
			"module":     module,
			"metricsets": []string{metricset},
			"hosts":      []string{server.URL},
		}

		f := mbtest.NewEventsFetcher(t, config)

		events, err := f.Fetch()
		assert.NoError(t, err)

		if update {
			eventsJSON, _ := json.MarshalIndent(events, "", "\t")
			err = ioutil.WriteFile(test.ExpectedFile, eventsJSON, 0644)
			assert.NoError(t, err)
		}

		// Read expected events from reference file
		expected, err := ioutil.ReadFile(test.ExpectedFile)
		if err != nil {
			t.Fatal(err)
		}

		var expectedEvents []common.MapStr
		err = json.Unmarshal(expected, &expectedEvents)
		if err != nil {
			t.Fatal(err)
		}

		for _, event := range events {
			// ensure the event is in expected list
			found := -1
			for i, expectedEvent := range expectedEvents {
				if event.String() == expectedEvent.String() {
					found = i
					break
				}
			}
			if found > -1 {
				expectedEvents = append(expectedEvents[:found], expectedEvents[found+1:]...)
			} else {
				t.Errorf("Event was not expected: %+v", event)
			}
		}

		if len(expectedEvents) > 0 {
			t.Error("Some events were missing:")
			for _, e := range expectedEvents {
				t.Error(e)
			}
			t.Fatal()
		}
	}
}
