package namesilo

import (
	"encoding/xml"
	"strconv"
)

// RegisterDomain attempts to register the given domain for the given years with any additionally provided options.  Options should be in "key=value" form.
func (c Client) RegisterDomain(d string, y int, opts ...string) (int, string, error) {
	params := make([]string, len(opts)+2)
	params[0] = "domain=" + d
	params[1] = "years=" + strconv.Itoa(y)
	params = append(params, opts...)

	body, err := c.callAPI("registerDomain", params...)
	if err != nil {
		return 0, "", err
	}

	var rdr registerDomainReply
	err = xml.Unmarshal(body, &rdr)
	if err != nil {
		return 0, "", err
	}
	return rdr.Code, rdr.OrderAmt, nil
}
