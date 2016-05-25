package namesilo

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	APIKey    string
	ReadLimit int64 // ReadLimit is the maximum number of bytes that can be read per request.
	version   uint8
	encoding  string
	server    string
}

// NewClient returns a client for the given API Key with a default read limit per request.
func NewClient(k string) Client {
	return Client{
		APIKey:    k,
		ReadLimit: 5 * 1024 * 1024, // Default 5MB
		version:   1,
		encoding:  "xml",
		server:    "https://www.namesilo.com",
	}
}

func (c Client) callAPI(op string, params ...string) ([]byte, error) {
	qs := make([]string, len(params)+3)
	qs[0] = "version=" + strconv.Itoa(int(c.version))
	qs[1] = "type=" + c.encoding
	qs[2] = "key=" + c.APIKey
	qs = append(qs, params...)

	r, err := http.Get(c.server + "/api/" + op + "?" + strings.Join(qs, "&"))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New("Namesilo guarantees HTTP Status Code 200. Received HTTP Status Code:" + strconv.Itoa(r.StatusCode) + ".")
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, c.ReadLimit))
	return body, err
}
