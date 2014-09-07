package versio

import (
	"net/url"
	"strconv"
)

// DomainsCreateContact adds a new contact. Optional parameters may be left empty (""), and other info available at API documentation. 
func DomainsCreateContact (companyname, initials, surname, email, phone, street, hnr, hnradd, zipcode, city, country string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsCreateContact")
	if len(companyname) > 0 {
		data.Add("companyname", companyname)
	}
	data.Add("initials", initials)
	data.Add("surname", surname)
	data.Add("email", email)
	data.Add("phone", phone)
	data.Add("street", street)
	data.Add("hnr", hnr)
	if len(hnradd) > 0 {
		data.Add("hnradd", hnradd)
	}
	data.Add("zipcode", zipcode)
	data.Add("city", city)
	data.Add("country", country)
	return Send(data)
}

// DomainsListContacts shows all contacts available
func DomainsListContacts () (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListContacts")
	return Send(data)
}

// DomainsListSingleContact shows a single contact
func DomainsListSingleContact (contact_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListSingleContact")
	data.Add("id", strconv.Itoa(contact_id))
	return Send(data)
}

// DomainsDeleteContact deletes a single contact from the list
func DomainsDeleteContact (contact_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDeleteContact")
	data.Add("id", strconv.Itoa(contact_id))
	return Send(data)
}

// DomainsChangeOwner changes the domain holder for a domain name. NOTE: this is not always free of charge! EPP-code is only for .BE domain-names. 
func DomainsChangeOwner (domain, tld string, newcontact_id int, eppcode string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsChangeOwner")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("newcontact_id", strconv.Itoa(newcontact_id))
	if len(eppcode) > 0 {
		data.Add("eppcode", eppcode)
	}
	return Send(data)
}

// DomainsListCountries lists the available countries along with country codes available. 
func DomainsListCountries () (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListCountries")
	return Send(data)
}
