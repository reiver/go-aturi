package aturi_test

import (
	"fmt"

	"github.com/reiver/go-aturi"
)

func ExampleValidatePrefix() {

	var uri string = "at:example.com"

	err := aturi.ValidatePrefix(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: aturi: AT-URI "at:example.com" is not valid because it does not have "//" after "at:"
}
