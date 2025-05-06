# go-aturi

Package **aturi** provides tools for working with AT-URIs, for the Go programming language.

AT-URIs are used by the Bluesky network and its AT-protocol, and are defined at:
https://atproto.com/specs/at-uri-scheme

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-aturi

[![GoDoc](https://godoc.org/github.com/reiver/go-aturi?status.svg)](https://godoc.org/github.com/reiver/go-aturi)

## Examples

Here are some examples for using package **aturi**.

### Validation Example

You can validate an AT-URI with code similar to the following:

```golang
import "github.com/reiver/go-aturi"

// ...

var atURI string = "at://did:plc:scewmn2pl3oz36mxme2b6czz"

// ...

err := aturi.Validate(atURI)
```

### Split Example

You can split an AT-URI into its components with code similar to the following:

```golang
import "github.com/reiver/go-aturi"

// ...

var atURI string = "at://did:plc:scewmn2pl3oz36mxme2b6czz"

// ...

authority, collection, rkey, query, fragment, err := aturi.Split(atURI)
```

## Join Example

You can create a normalized AT-URI with code similar to the following:

```golang
import "github.com/reiver/go-aturi"

// ...

var authority  string = "did:plc:scewmn2pl3oz36mxme2b6czz"
var collection string = "com.example.fooBar"
var rkey       string = "wxyz1234"
var query      string = ""
var fragment   string = ""

var atURI string = aturi.Join(authority, collection, rkey, query, fragment)
```

## Import

To import package **aturi** use `import` code like the follownig:
```
import "github.com/reiver/go-aturi"
```

## Installation

To install package **aturi** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-aturi
```

## Author

Package **aturi** was written by [Charles Iliya Krempeaux](http://reiver.link)

## See Also

* https://github.com/reiver/go-athandle
* https://github.com/reiver/go-atproto
* https://github.com/reiver/go-aturi
* https://github.com/reiver/go-bsky
* https://github.com/reiver/go-did
* https://github.com/reiver/go-didplc
* https://github.com/reiver/go-nsid
* https://github.com/reiver/go-xrpc
* https://github.com/reiver/go-xrpcuri
