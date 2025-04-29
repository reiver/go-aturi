package aturi_test

import (
	"fmt"

	"github.com/reiver/go-aturi"
)

func ExampleNormalizeCollection() {

	var authority string = "COM.Example.fooBar"

	normalizedCollection := aturi.NormalizeCollection(authority)

	fmt.Printf("original collection:   %s\n", authority)
	fmt.Printf("normalized collection: %s\n", normalizedCollection)

	// Output:
	// original collection:   COM.Example.fooBar
	// normalized collection: com.example.fooBar
}
