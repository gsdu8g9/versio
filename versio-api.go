package versio

import (
	"fmt"
	"strings"
	"net/url"
	"bytes"
	"mime/multipart"
)

var klantNr 	= "64349"
var klantPass 	= "d4f041bea55ae1524318386b607995112589983d"
var sandBox 	= "1"
var apiUrl		= "https://www.secure.versio.nl/api/api_server.php"

// Initialize sets the required variables (i.e. Password) for the Versio-API authentication. The password has to be SHA1-hashed already.
func Initialize(nr, pw string, sand bool) {
	klantNr = nr
	klantPass = pw
	if sand {
		sandBox = "1"
	} else {
		sandBox = "0"
	}
}

func search(domain, tld string) {

	// Set POST data
	data := url.Values{}
	data.Set("command",	"DomainsCheckAvailability")
	data.Set("domain", 	domain)
	data.Set("tld",		tld)
	data.Set("klantPw", klantPass)
	data.Set("klantId", klantNr)
	data.Set("sandBox", sandBox)

	// Write POST-data to Multipart/Form-Data
	Mb := &bytes.Buffer{}
	Mw := multipart.NewWriter(Mb)
	for key, val := range data {
		err := Mw.WriteField(key, val[0])
		if err != nil {
			fmt.Println(&Mb, &Mw, err)
		}
	} 
	Mw.Close()
	boundary := Mw.Boundary()
	
	
	// Prepare client and request
	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiUrl, Mb)
	r.Close = true
	r.Header.Set("Expect", "100-continue")
	r.Header.Set("Content-Type", "multipart/form-data; boundary=" + boundary)
	
	// Do request and parse results
	resp, _ := client.Do(r)
	contents, _ := ioutil.ReadAll(resp.Body)
	mapped := contentToMap(contents)
	
	// Return results to channel
	if mapped["success"] == "0" {
	dbg.Print(dbg.ERROR, time.Now().String(), "Domain-check", domain, tld, mapped)
	}
	avail := mapped["success"] == "1" && mapped["status"] == "1"
	result := &searchResult{domain + "." + tld, avail}
	cache[domain + "." + tld] = searchQuery{time.Now(), avail}
	
	}

func contentToMap(content []byte) map[string]string {
	returnable := make(map[string]string)
	splitLines := strings.Split(string(content), "\n")
	for _, line := range splitLines {
		splitPair := strings.Split(line, "=")
		returnable[splitPair[0]] = splitPair[1]
	}
	return returnable
}
