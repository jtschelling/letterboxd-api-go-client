# letterboxd-api-go-client

[![CI](https://github.com/jtschelling/letterboxd-api-go-client/workflows/test/badge.svg)](https://github.com/jtschelling/letterboxd-api-go-client/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jtschelling/letterboxd-api-go-client)](https://goreportcard.com/report/github.com/jtschelling/letterboxd-api-go-client)
[![GoDoc](https://godoc.org/github.com/jtschelling/letterboxd-api-go-client?status.svg)](https://godoc.org/github.com/jtschelling/letterboxd-api-go-client)
[![License](https://img.shields.io/github/license/jtschelling/letterboxd-api-go-client)](https://opensource.org/licenses/MIT)

## API Access

The Letterboxd API is in [private beta](https://letterboxd.com/api-beta/). Access can be requested by contacting the Letterboxd team at api@letterboxd.com

## Installation

## Usage

### Request Signing

Create a salted string of the form `[METHOD]\u0000[URL]\u0000[BODY]` where [METHOD] is GET, POST, etc., [URL] is the fully-qualified request URL including the `apikey, nonce, timestamp` and any other method parameters, and [BODY] is a JSON-encoded string (for POST, PATCH and DELETE requests) or empty (for GET requests).

This computation is done in the request builder.

Example:

```{Golang}
METHOD = "GET"
URL = "https://api.letterboxd.com/api/v0/auth/token?apikey=FAKE_KEY&nonce=123-456-789&timestamp=1607292265"
BODY = ""
```


## Contributing

Issues are tracked on Github and can be submitted [here](https://github.com/jtschelling/letterboxd-api-go-client/issues/new).
Pull requests are welcome and will be reviewed in a timely manner.
