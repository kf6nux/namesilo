package namesilo

import "encoding/xml"

// GetDomainInfo returns namesilo.DomainInfo for the given domain.
func (c Client) GetDomainInfo(d string) (*DomainInfo, error) {
	body, err := c.callAPI("getDomainInfo", "domain="+d)
	if err != nil {
		return nil, err
	}

	var gdr getDomainReply
	err = xml.Unmarshal(body, &gdr)
	if err != nil {
		return nil, err
	}
	if gdr.Code != 300 {
		return nil, gdr.StdReply
	}
	return &gdr.DomainInfo, nil
}
