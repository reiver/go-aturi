package aturi_test

import (
	"fmt"

	"github.com/reiver/go-aturi"
)

func ExampleSplit() {

	var uri string = "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y"

	authority, collection, rkey, query, fragment, err := aturi.Split(uri)
	if nil != err {
		fmt.Printf("ERROR: problem splitting AT-URI: %s\n", err)
		return
	}

	fmt.Printf("authority:  %q\n", authority)
	fmt.Printf("collection: %q\n", collection)
	fmt.Printf("rkey:       %q\n", rkey)
	fmt.Printf("query:      %q\n", query)
	fmt.Printf("fragment:   %q\n", fragment)

	// Output:
	// authority:  "did:plc:scewmn2pl3oz36mxme2b6czz"
	// collection: "com.example.foorBar"
	// rkey:       "3jui7kd54zh2y"
	// query:      ""
	// fragment:   ""
}
