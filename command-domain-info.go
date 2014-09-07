package versio

import "net/url"

// DomainsTLDInfo retrieves a list of available extensions (pricing included). If tld is nil, all extensions will be checked.
func DomainsTLDInfo (tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsTLDInfo")
	if len(tld) > 0 {
		data.Add("tld", tld)
	}
	return Send(data)
}

// DomainsCheckAvailability check whether or not the domain name is available. Limited usage: 100 + 8 * registered-domains. 
func DomainsCheckAvailability(domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsCheckAvailability")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsListActive gives a list of active domain names, with expiration date, auto-renew setting and nameservers. 
func DomainsListActive() (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListActive")
	return Send(data)
}

// DomainsListInactive gives a list of inactive domain names, with expiration date, auto-renew setting and nameservers. 
func DomainsListInactive() (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListInactive")
	return Send(data)
}

// DomainsListSingle gives information for one domain name, such as expiration date, auto-renew setting and nameservers. 
func DomainsListSingle(domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListSingle")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsDelete deletes / declaims the domain name.
func DomainsDelete(domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDelete")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsListTransfers shows a list of ongoing domain transfers, with information such as transfer status, text, and update time.
func DomainsListTransfers() (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListTransfers")
	return Send(data)
}

// DomainsListTransferredAway shows a list domains that were transferred in the past 30 days, including a timestamp of each transfer.
func DomainsListTransferredAway() (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListTransferredAway")
	return Send(data)
}

// DomainsGetToken returns the transfer-token for a domain name (ONLY for .NL domain names)
func DomainsGetToken(domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsGetToken")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsSendEPP_be sends the EPP-code for a .BE extension to the domain holder. This domain does NOT have to be in your account to do this. 
func DomainsSendEPP_be(domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsSendEPP_be")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsSendEPP sends the EPP-code for a .COM/NET/INFO/BIZ/ORG extension to the administrator of the domain (WHOIS). This domain does NOT have to be in your account to do this. 
func DomainsSendEPP(domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsSendEPP")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}

// DomainsGetEPP returns the EPP token for .COM/NET?INFO/BIZ/ORG extensions. May be unavailable - in that case: use  DomainsSendEPP.
func DomainsGetEPP(domain, tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsGetEPP")
	data.Add("domain", domain)
	data.Add("tld", tld)
	return Send(data)
}
