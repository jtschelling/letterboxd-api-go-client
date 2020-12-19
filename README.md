# letterboxd-api-go-client

[![CI](https://github.com/jtschelling/letterboxd-api-go-client/workflows/test/badge.svg)](https://github.com/jtschelling/letterboxd-api-go-client/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jtschelling/letterboxd-api-go-client)](https://goreportcard.com/report/github.com/jtschelling/letterboxd-api-go-client)
[![GoDoc](https://pkg.go.dev/github.com/jtschelling/letterboxd-api-go-client?status.svg)](https://pkg.go.dev/github.com/jtschelling/letterboxd-api-go-client/letterboxd)
[![License](https://img.shields.io/github/license/jtschelling/letterboxd-api-go-client)](https://opensource.org/licenses/MIT)

## API Access

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
package main

import (
    "fmt"
    l "github.com/jtschelling/letterboxd-api-go-client/letterboxd"
)

func main() {
    c := l.NewClientFromKey("API_KEY", "API_SECRET")
    t, _ := c.GetUserAccessToken("USERNAME", "PASSWORD")
    m, _ := c.GetSelf(t)

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

Relevant test updates are required. CI is run on pull request and subsequent commits with Github Actions.
