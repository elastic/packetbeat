package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/conditions"
	"github.com/stretchr/testify/require"
)

func TestCheckJsonExpression(t *testing.T) {
	simpleJson := "{\"foo\": \"hi\", \"bar\": 3, \"baz\": {\"bot\": \"blah\"}}"
	arrayJson := fmt.Sprintf("[%s, %s]", simpleJson, simpleJson)
	var tests = []struct {
		description   string
		body          string
		expression    string
		expectSuccess bool
	}{
		{
			"good match succeeds",
			simpleJson,
			"foo == \"hi\" && bar == 3",
			true,
		},
		{
			"bad match fails",
			simpleJson,
			"foo == \"hi\" && bar == 1000",
			false,
		},
		{
			"deep match succeeds",
			simpleJson,
			"baz == {\"bot\": \"blah\"}",
			true,
		},
		{
			"bad deep match fails",
			simpleJson,
			"baz == {\"bot\": \"nope\"}",
			false,
		},
		{
			"good match succeeds with jsonpath",
			simpleJson,
			"$.baz.bot == \"blah\"",
			true,
		},
		{
			"bad match fails with jsonpath",
			simpleJson,
			"$.baz.bot == \"nope\"",
			false,
		},
		{
			"good array json matches jsonpath",
			arrayJson,
			"$[0].baz.bot == \"blah\"",
			true,
		},
		{
			"bad array json matches jsonpath",
			arrayJson,
			"$[0].baz.bot == \"nope\"",
			false,
		},
	}

	for _, test := range tests[6:] {
		t.Run(test.description, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, test.body)
			}))
			defer ts.Close()

			res, err := http.Get(ts.URL)
			if err != nil {
				log.Fatal(err)
			}

			checker, err := checkJSON(
				[]*jsonResponseCheck{
					{
						Description: test.description,
						Expression:  test.expression,
					},
				},
			)

			require.NoError(t, err)
			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			checkRes := checker(res, string(body))

			if result := checkRes == nil; result != test.expectSuccess {
				if test.expectSuccess {
					t.Fatalf("Expected expression: '%s' to match body: %s. got: %s", test.expression, test.body, checkRes)
				} else {
					t.Fatalf("Did not expect expression: '%s' to match body: %s. got: %s", test.expression, test.body, checkRes)
				}
			}
		})
	}

}

func TestCheckJsonCondition(t *testing.T) {
	fooBazEqualsBar := common.MustNewConfigFrom(map[string]interface{}{"equals": map[string]interface{}{"foo": map[string]interface{}{"baz": "bar"}}})
	fooBazEqualsBarConf := &conditions.Config{}
	err := fooBazEqualsBar.Unpack(fooBazEqualsBarConf)
	require.NoError(t, err)

	fooBazEqualsBarInt := common.MustNewConfigFrom(map[string]interface{}{"equals": map[string]interface{}{"foo": 1}})
	fooBazEqualsBarIntConf := &conditions.Config{}
	err = fooBazEqualsBarInt.Unpack(fooBazEqualsBarIntConf)
	require.NoError(t, err)
	fooBazEqualsBarIntDesc := "foo equals 1"

	fooBazEqualsBarDesc := "foo.baz equals bar"

	var tests = []struct {
		description string
		body        string
		condDesc    string
		condConf    *conditions.Config
		result      bool
	}{
		{
			"positive match",
			"{\"foo\": {\"baz\": \"bar\"}}",
			fooBazEqualsBarDesc,
			fooBazEqualsBarConf,
			true,
		},
		{
			"Negative match",
			"{\"foo\": 123}",
			fooBazEqualsBarDesc,
			fooBazEqualsBarConf,
			false,
		},
		{
			"unparseable",
			`notjson`,
			fooBazEqualsBarDesc,
			fooBazEqualsBarConf,
			false,
		},
		{
			"positive int match",
			"{\"foo\": 1}",
			fooBazEqualsBarIntDesc,
			fooBazEqualsBarIntConf,
			true,
		},
		{
			"Negative int match",
			"{\"foo\": 2}",
			fooBazEqualsBarIntDesc,
			fooBazEqualsBarIntConf,
			false,
		},
		{
			"Negative string match against int",
			"{\"foo\": \"some string\"}",
			fooBazEqualsBarIntDesc,
			fooBazEqualsBarIntConf,
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, test.body)
			}))
			defer ts.Close()

			res, err := http.Get(ts.URL)
			if err != nil {
				log.Fatal(err)
			}

			checker, err := checkJSON([]*jsonResponseCheck{{Description: test.condDesc, Condition: test.condConf}})
			require.NoError(t, err)
			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			checkRes := checker(res, string(body))

			if result := checkRes == nil; result != test.result {
				if test.result {
					t.Fatalf("Expected condition: '%s' to match body: %s. got: %s", test.condDesc, test.body, checkRes)
				} else {
					t.Fatalf("Did not expect condition: '%s' to match body: %s. got: %s", test.condDesc, test.body, checkRes)
				}
			}
		})
	}
}
