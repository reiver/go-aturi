package aturi_test

import (
        "fmt"

        "github.com/reiver/go-aturi"
)

func ExampleValidateAuthority() {

        var authority string = "user:pass@archive.org"

        err := aturi.ValidateAuthority(authority)

        fmt.Printf("authority: %s\n", authority)
        fmt.Printf("validation error: %s\n", err)

        // Output:
        // authority: user:pass@archive.org
        // validation error: aturi: authority may not have an "@" in it
}
