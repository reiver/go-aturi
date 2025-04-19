# go-aturi

Package **aturi** provides tools for working with AT-URIs, for the Go programming language.

AT-URIs are used by the Bluesky network and its AT-protocol.

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

var uri string = "at://did:plc:scewmn2pl3oz36mxme2b6czz"

// ...

err := aturi.Validate(uri)
```

### Split Example

You can split an AT-URI into its components with code similar to the following:

```golang
import "github.com/reiver/go-aturi"

// ...

var uri string = "at://did:plc:scewmn2pl3oz36mxme2b6czz"

// ...

authority, collection, rkey, query, fragment, err := aturi.Split(uri)
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
