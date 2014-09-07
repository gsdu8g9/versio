package versio

import (
	"net/url"
	"strconv"
)

// DomainsListCategories returns a list of domain categories
func DomainsListCategories () (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsListCategories")
	return Send(data)
}

// DomainsCreateCategory creates a new (empty) DNS category with the given name
func DomainsCreateCategory (category_name string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsCreateCategory")
	data.Add("name", category_name)
	return Send(data)
}

// DomainsDeleteCategory removes a domain-category from the list
func DomainsDeleteCategory (category_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsDeleteCategory")
	data.Add("id", strconv.Itoa(category_id))
	return Send(data)
}

// DomainsChangeCategory removes a domain-category from the list
func DomainsChangeCategory (domain, tld string, category_id int) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsChangeCategory")
	data.Add("domain", domain)
	data.Add("tld", tld)
	data.Add("category_id", strconv.Itoa(category_id))
	return Send(data)
}
