package versio

import (
	"net/url"
	"strconv"
)

// DomainsChangeDNSTemplate changes the DNS template of the domain
func DomainsChangeDNSTemplate (domain, tld string, template int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsChangeDNSTemplate")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("template", strconv.Itoa(template))
	return Send(data)
}

// DomainsDNSon turns on DNS administration for domain
func DomainsDNSon (domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSon")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsDNSOff turns off DNS administration for domain
func DomainsDNSOff (domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSOff")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsDNSListRecords lists the DNS records for the domain
func DomainsDNSListRecords (domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSListRecords")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsDNSListRedirections returns a list of redirections for the domain
func DomainsDNSListRedirections (domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSListRedirections")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}


// DomainsDNSAddRecord adds a new DNS record for the domain. Any prio-value <= 0 will omit the optional prio-parameter. 
func DomainsDNSAddRecord (domain, tld, name string, recordType RecordType, value string, prio int, ttl TTLValue) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSAddRecord")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("name", name)
	data.Add("type", recordType.String())
	data.add("value", value)
	if prio > 0 {
		data.add("prio", strconv.Itoa(prio))
	}
	data.add("ttl", ttl.String())
	return Send(data)
}

// DomainsDNSAddRedirectionURL changes the DNS template of the domain, using a url.URL as parameter value
func DomainsDNSAddRedirectionURL (domain, tld, from string, to url.URL) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSAddRedirection")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("from", from)
	data.Add("url", to.String())
	return Send(data)
}

// DomainsDNSAddRedirection changes the DNS template of the domain, using a string as parameter value
func DomainsDNSAddRedirection (domain, tld, from, to string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSAddRedirection")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("from", from)
	data.Add("url", to)
	return Send(data)
}

// DomainsDNSDeleteRecord deletes a DNS record for a domain
func DomainsDNSDeleteRecord (domain, tld string, id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSDeleteRecord")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("id", strconv.Itoa(id))
	return Send(data)
}

// DomainsDNSDeleteRedirection changes the DNS template of the domain
func DomainsDNSDeleteRedirection (domain, tld string, id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDNSDeleteRedirection")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("id", strconv.Itoa(id))
	return Send(data)
}
