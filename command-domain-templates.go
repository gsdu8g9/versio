package versio

import (
	"net/url"
	"strconv"
)

// DNSTemplatesList returns a list of DNS templates
func DNSTemplatesList () (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesList")
	return Send(data)
}

// DNSTemplatesListRecords returns the DNS records of a selected DNS template
func DNSTemplatesListRecords (template_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesListRecords")
	data.Add("id", strconv.Itoa(template_id))
	return Send(data)
}

// DNSTemplatesListRedirections shows a list of DNS redirections of a selected DNS template
func DNSTemplatesListRedirections (template_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesListRedirections")
	data.Add("id", strconv.Itoa(template_id))
	return Send(data)
}

// DNSTemplatesCreate creates a new (empty) DNS template with given name
func DNSTemplatesCreate (template_name string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesCreate")
	data.Add("name", template_name)
	return Send(data)
}

// DNSTemplateAddRecord adds a new DNS record for the given template. Any prio-value <= 0 will omit the optional prio-parameter. 
func DNSTemplateAddRecord (template_id int, name string, rtype RecordType, value string, prio int, ttl TTLValue) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplateAddRecord")
	data.Add("id", strconv.Itoa(template_id))
	data.Add("name", name)
	data.Add("type", rtype.String())
	data.Add("value", value)
	if prio > 0 {
		data.Add("prio", strconv.Itoa(prio))
	}
	data.Add("ttl", ttl.String())
	return Send(data)
}

// DNSTemplatesAddRedirection adds a new redirect-record for the given DNS template
func DNSTemplatesAddRedirection (template_id int, from string, to url.URL) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesAddRedirection")
	data.Add("id", strconv.Itoa(template_id))
	data.Add("from", from)
	data.Add("url", to.String())
	return Send(data)
}

// DNSTemplatesDeleteRecord removes a DNS record from a DNS template, DNSListRecordsFromTemplate returns the required Record-ID. 
func DNSTemplatesDeleteRecord (record_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesDeleteRecord")
	data.Add("id", strconv.Itoa(record_id))
	return Send(data)
}

// DNSTemplatesDeleteRedirection removes a DNS-redirect-record from a DNS template, DNSListRecordsFromTemplate returns the required Record-ID. 
func DNSTemplatesDeleteRedirection (record_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesDeleteRedirection")
	data.Add("id", strconv.Itoa(record_id))
	return Send(data)
}

// DNSTemplatesDelete removes a template from the list of templates. 
func DNSTemplatesDelete (template_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DNSTemplatesDelete")
	data.Add("id", strconv.Itoa(template_id))
	return Send(data)
}
