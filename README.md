# letterboxd-api-go-client

[![CI](https://github.com/jtschelling/letterboxd-api-go-client/workflows/test/badge.svg)](https://github.com/jtschelling/letterboxd-api-go-client/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jtschelling/letterboxd-api-go-client)](https://goreportcard.com/report/github.com/jtschelling/letterboxd-api-go-client)
[![GoDoc](https://pkg.go.dev/github.com/jtschelling/letterboxd-api-go-client?status.svg)](https://pkg.go.dev/github.com/jtschelling/letterboxd-api-go-client/letterboxd)
[![License](https://img.shields.io/github/license/jtschelling/letterboxd-api-go-client)](https://opensource.org/licenses/MIT)

## API Access

[API Reference Page](http://api-docs.letterboxd.com/)

The Letterboxd API is in [private beta](https://letterboxd.com/api-beta/). Access can be requested by contacting the Letterboxd team at api@letterboxd.com

## Installation

`go get github.com/jtschelling/letterboxd-api-go-client`

```{golang}
package main

import (
    l "github.com/jtschelling/letterboxd-api-go-client/letterboxd
)
```

## Usage

Refer to the [GoDoc](https://pkg.go.dev/github.com/jtschelling/letterboxd-api-go-client/letterboxd) for more usage information

```{Golang}
// main.go
package main

import (
    "fmt"
    l "github.com/jtschelling/letterboxd-api-go-client/letterboxd"
)

func main() {
    // Create a client object. Requires API keys
    c := l.NewClientFromKey("API_KEY", "API_SECRET")

    // Retrieve an user session token to provide with most further API requests
    t, _ := c.GetAuthToken("USERNAME", "PASSWORD")

    // Simple request to get information on currently authorized user
    m, _ := c.GetMe(t)

    fmt.Printf("Auth Token: %s\nMember ID: %s\n", t, m.Member.ID)
}
```

```{Bash}
$ go get github.com/jtschelling/letterboxd-api-go-client
$ go run main.go
Auth Token: ABC
Member ID: 123
```

### Request Signing

Create a salted string of the form `[METHOD]\u0000[URL]\u0000[BODY]` where [METHOD] is GET, POST, etc., [URL] is the fully-qualified request URL including the `apikey, nonce, timestamp` and any other method parameters, and [BODY] is a JSON-encoded string (for POST, PATCH and DELETE requests) or empty (for GET requests).

This computation is done in the request builder.

## Contributing

Issues are tracked on Github and can be submitted [here](https://github.com/jtschelling/letterboxd-api-go-client/issues/new).
Pull requests are welcome and will be reviewed in a timely manner.

Please review [CONTRIBUTING.md](CONTRIBUTING.md) before submitting a pull request.
