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
```

### One time: initialization
```go
versio.Initialize("12345", "de9f2c7fd25e1b3afad3e85a0bd17d9b100db4b3", false)
```

### Per query: either one of the following:
#### Use 'raw' Send-function
```go
// Prepare query data
data := url.Values{}
data.Set("command",	    "DomainsCheckAvailability")
data.Set("domain", 	    "example")
data.Set("tld",	        "com")

// Actually making request
mapped, err := versio.Send(data)
if err != nil {
	// error handling ...
}

// using mapped["success"] and others ...
```

#### Use implemented command
```go
mapped, err := versio.DomainsCheckAvailability("domain", "tld")
if err != nil { 
	// error handling .
}

// using mapped["success"] and others ...
```

Documentation
-------------
Documentation of the Golang implementation is available at http://godoc.org/github.com/HuperWebs/versio

Documentation of the Versio API is available in the PHP-version download at http://www.versio.nl/api.php
