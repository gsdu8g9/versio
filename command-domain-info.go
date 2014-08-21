package versio

import "net/url"

import (
)

/** DomainsTLDInfo retrieves a list of available extensions (pricing included). If tld is nil, all extensions will be checked.
 *
 * Returns map containing: 
 *     success 					(whether or not the command was successful)
 *     command_repsonse_message 		(error message, if success == 0)
 *     total_count 					(amount of extensions, if success == 1)
 *     tld_X 						(extension beloning to the information, range X from 1 to total_count, if success == 1)
 *     newprice_X 					(price for extension X for the first year, if success == 1)
 *     renewprice_X					(price for extension X for successive years, if success == 1)
 *     minimumyears_X					(minimum amount of years the domain has to be registered, if success == 1)
 */
func DomainsTLDInfo (tld string) (map[string]string, error) {
	data := url.Values{}
	data.Add("command", "DomainsTLDInfo")
	if tld != nil {
		data.Add("tld", tld)
	}
	return Send(data)
}
