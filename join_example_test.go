package aturi_test

import (
	"fmt"

	"github.com/reiver/go-aturi"
)

func ExampleJoin() {

	var authority  string = "did:plc:scewmn2pl3oz36mxme2b6czz"
	var collection string = "com.example.foorBar"
	var rkey       string = "3jui7kd54zh2y"

	uri := aturi.Join(authority, collection, rkey, "", "")

	fmt.Printf("uri: %q\n", uri)

	// Output:
	// uri: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y"
}
