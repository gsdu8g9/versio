// Package versio allows access to the Versio-API. 
package versio

import (
	"net/http"
	"strings"
	"net/url"
	"bytes"
	"mime/multipart"
	"io/ioutil"
	"errors"
)

var klantNr 	string
var klantPass 	string
var sandBox 	string
var initialized bool
var apiUrl		= "https://www.secure.versio.nl/api/api_server.php"

// Initialize sets the required variables (i.e. Password) for the Versio-API authentication. The password has to be SHA1-hashed already.
func Initialize(ClientNumber, ClientPassword string, SandboxMode bool) {
	initialized = true
	klantNr = ClientNumber
	klantPass = ClientPassword
	if SandboxMode {
		sandBox = "1"
	} else {
		sandBox = "0"
	}
}

// Send with channel is similar to 'Send', but sends the return-value to the given channel instead. Panics on error. 
func SendWithChannel(CommandParameters url.Values, ch chan *map[string]string) {
	output, err := Send(CommandParameters)
	if err != nil {
		panic(err)
	} else {
		ch <- output
	}
}

// Send creates a POST-request for the given values to the API server. Returns the response as a map with keys such as 'success'. 
func Send(CommandParameters url.Values) (*map[string]string, error) {
	if !initialized {
		return nil, errors.New("Must call versio.Initialize first. ")
	}
	// Set POST data
	data := url.Values{}
	data.Set("klantPw", klantPass)
	data.Set("klantId", klantNr)
	data.Set("sandBox", sandBox)
	
	// Write POST-data to Multipart/Form-Data
	Mb := &bytes.Buffer{}
	Mw := multipart.NewWriter(Mb)
	for key, val := range data {
		Mw.WriteField(key, val[0])
	}
	for key, val := range CommandParameters {
		Mw.WriteField(key, val[0])
	}
	Mw.Close()
	// Save boundary for Content-Type
	boundary := Mw.Boundary()
	
	// Prepare client and request
	client := &http.Client{}
	r, err := http.NewRequest("POST", apiUrl, Mb)
	if err != nil {
		return nil, err
	}
	r.Close = true	// Connection: Close
	r.Header.Set("Expect", "100-continue")
	r.Header.Set("Content-Type", "multipart/form-data; boundary=" + boundary)
	
	// Do request and parse results
	if resp, err := client.Do(r); err != nil {
		return nil, err
	} else {
		if contents, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else {
			mapped := contentToMap(contents)
			return &mapped, nil
		}
	}
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
