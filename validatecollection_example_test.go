package aturi_test

import (
        "fmt"

        "github.com/reiver/go-aturi"
)

func ExampleValidateCollection() {

        var collection string = "COM.Example.fooBar"

        err := aturi.ValidateCollection(collection)

        fmt.Printf("collection (i.e., NSID): %s\n", collection)
        fmt.Printf("validation error: %s\n", err)

        // Output:
	// collection (i.e., NSID): COM.Example.fooBar
        // validation error: nsid: character №0 ('C') (U+0043) of domain-authority part №0 ("COM") of domain-authority ("COM.Example") of nsid ("COM.Example.fooBar") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')
}
