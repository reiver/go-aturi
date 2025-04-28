package aturi_test

import (
        "fmt"

        "github.com/reiver/go-aturi"
)

func ExampleValidate() {

        var atURI string = "at://archive.org/COM.Example.fooBar/3jui7kd54zh2y"

        err := aturi.Validate(atURI)

        fmt.Printf("AT-URI: %s\n", atURI)
        fmt.Printf("validation error: %s\n", err)

        // Output:
        // AT-URI: at://archive.org/COM.Example.fooBar/3jui7kd54zh2y
        // validation error: aturi: AT-URI "at://archive.org/COM.Example.fooBar/3jui7kd54zh2y" has a collection "COM.Example.fooBar" that is not a valid NSID: nsid: character №0 ('C') (U+0043) of domain-authority part №0 ("COM") of domain-authority ("COM.Example") of nsid ("COM.Example.fooBar") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')
}
