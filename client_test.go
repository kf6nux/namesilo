package namesilo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("myapikey")
	if c.APIKey != "myapikey" ||
		c.ReadLimit == 0 ||
		c.version != 1 ||
		c.encoding != "xml" ||
		c.server != "https://www.namesilo.com" {
		t.Fatal("NewClient initilized client incorrectly.")
	}
}

func TestCallAPI(t *testing.T) {
	c := NewClient("myapikey")
	op := "someEndpoint"
	extraParams := make(map[string]string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/"+op {
			t.Error("callAPI should")
		}

		v := r.URL.Query()
		i, _ := strconv.Atoi(v.Get("version"))
		if i != int(c.version) {
			t.Error("callAPI should send version.")
		}
		if c.encoding != v.Get("type") {
			t.Error("callAPI should send type.")
		}
		if c.APIKey != v.Get("key") {
			t.Error("callAPI should send api key.")
		}
		for k, kv := range extraParams {
			if v.Get(k) != kv {
				t.Error("callAPI should send any extra key-value pairs in the query string.")
			}
		}

		w.WriteHeader(200)
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	c.server = ts.URL
	// test without extra params
	c.callAPI(op)

	// test with extra params
	op = "otherEndpoint"
	extraParams["extraKey"] = "extraValue"
	c.callAPI(op, "extraKey=extraValue")
}
