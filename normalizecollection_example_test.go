package aturi_test

import (
	"fmt"

	"github.com/reiver/go-aturi"
)

func ExampleNormalizeCollection() {

	var collection string = "COM.Example.fooBar"

	normalizedCollection := aturi.NormalizeCollection(collection)

	fmt.Printf("original collection:   %s\n", collection)
	fmt.Printf("normalized collection: %s\n", normalizedCollection)

	// Output:
	// original collection:   COM.Example.fooBar
	// normalized collection: com.example.fooBar
}
