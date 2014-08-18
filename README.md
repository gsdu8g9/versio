Go-Versio-API
=============

The API for using the Versio ( http://www.versio.nl/ ) interface using the language Go. An account at Versio is required. 

Usage
-------------
    
```go
import (
    "net/url"
    "github.com/HuperWebs/versio"
)


// One time: initialization
versio.Initialize("12345", "de9f2c7fd25e1b3afad3e85a0bd17d9b100db4b3", false)

// Per-query: set information
data := url.Values{}
data.Set("command",	    "DomainsCheckAvailability")
data.Set("domain", 	    "example")
data.Set("tld",	        "com")

// Actually making request
mapped, error := versio.Send(data)
if error != nil {
	// error handling ...
}

// using mapped["success"] and others ...
```

Documentation
-------------
Documentation is available at http://godoc.org/github.com/HuperWebs/versio
