package namesilo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRegisterDomain(t *testing.T) {
	domain := "foo.com"
	years := 3
	code := 300
	charge := "6.42"
	extraParams := make(map[string]string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		i, _ := strconv.Atoi(v.Get("years"))
		if i != years {
			t.Error("RegisterDomain should send 'years' key.")
		}
		if domain != v.Get("domain") {
			t.Error("RegisterDomain should send 'domain' key.")
		}
		for k, kv := range extraParams {
			if v.Get(k) != kv {
				t.Error("RegisterDomain should send any extra key-value pairs in the query string.")
			}
		}
		w.WriteHeader(200)
		fmt.Fprintln(w, `<namesilo>
                            <request>
                                <operation>registerDomain</operation>
                                <ip>55.555.55.55</ip>
                            </request>
                            <reply>
                                <code>`+strconv.Itoa(code)+`</code>
                                <detail>success</detail>
                                <message>Your domain registration was successfully processed.</message>
                                <domain>namesilo.com</domain>
                                <order_amount>`+charge+`</order_amount>
                            </reply>
                        </namesilo>`)
	}))
	defer ts.Close()

	c := NewClient("myapikey")
	c.server = ts.URL

	// test without extra params
	rcode, rcharge, err := c.RegisterDomain(domain, years)
	if err != nil {
		t.Error("RegisterDomain should not return error on success. Error: " + err.Error())
	}
	if rcode != code {
		t.Error("RegisterDomain should return the status code in the XML.")
	}
	if rcharge != charge {
		t.Error("RegisterDomain should return the order_amount in the XML.")
	}

	// test with extra params
	extraParams["extraKey"] = "extraValue"
	c.RegisterDomain(domain, years, "extraKey=extraValue")
}
