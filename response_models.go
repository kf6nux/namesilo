package namesilo

import "strconv"

// StdReply may be be used as the error value of functions in this package.
type StdReply struct {
	Code   int    `xml:"reply>code"`
	Detail string `xml:"reply>detail"`
}

func (r StdReply) Error() string {
	return strconv.Itoa(r.Code) + ":" + r.Detail
}

type listDomainsReply struct {
	StdReply
	Domains []string `xml:"reply>domains>domain"`
}

type getDomainReply struct {
	StdReply
	DomainInfo
}

// DomainInfo is a direct mapping of namesilo's example response for getDomainInfo
type DomainInfo struct {
	Created          string   `xml:"reply>created"`
	Expires          string   `xml:"reply>expires"`
	Status           string   `xml:"reply>status"`
	Locked           string   `xml:"reply>locked"`
	Private          string   `xml:"reply>private"`
	AutoRenew        string   `xml:"reply>auto_renew"`
	TrafficType      string   `xml:"reply>traffic_type"`
	ForwardUrl       string   `xml:"reply>forward_url"`
	ForwardType      string   `xml:"reply>forward_type"`
	Nameservers      []string `xml:"reply>nameservers>nameserver"`
	RegistrantID     string   `xml:"reply>contact_ids>registrant"`
	AdministrativeID string   `xml:"reply>contact_ids>administrative"`
	TechnicalID      string   `xml:"reply>contact_ids>technical"`
	BillingID        string   `xml:"reply>contact_ids>billing"`
}

type registerDomainReply struct {
	StdReply
	OrderAmt string `xml:"reply>order_amount"`
}
