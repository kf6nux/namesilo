package namesilo

import "encoding/xml"

// ListDomains lists domains for a given portfolio ID.  Portfolio ID should be an empty string for the default behavior.
func (c Client) ListDomains(p string) ([]string, error) {
	var qs []string
	if len(p) > 0 {
		qs = []string{"portfolio=" + p}
	}

	body, err := c.callAPI("listDomains", qs...)
	if err != nil {
		return nil, err
	}

	var ldr listDomainsReply
	err = xml.Unmarshal(body, &ldr)
	if err != nil {
		return nil, err
	}
	if ldr.Code != 300 {
		return nil, ldr.StdReply
	}
	return ldr.Domains, nil
}
