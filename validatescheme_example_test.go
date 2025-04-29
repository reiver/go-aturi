package aturi_test

import (
	"fmt"

	"github.com/reiver/go-aturi"
)

func ExampleValidateScheme() {

	var uri string = "http://example.com/once/twice/thrice/fource.html"

	err := aturi.ValidateScheme(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: aturi: URI "http://example.com/once/twice/thrice/fource.html" is not an AT-URI because it does not begin with "at:"
}
