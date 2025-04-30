package aturi_test

import (
	"fmt"

	"github.com/reiver/go-aturi"
)

func ExampleSanitizeAuthority() {

	var authority string = "JoeBlow:Pass123@Example.COM"

	sanitizedAuthority := aturi.SanitizeAuthority(authority)

	fmt.Printf("original authority: %s\n", authority)
	fmt.Printf("sanitized authority:                %s\n", sanitizedAuthority)

	// Output:
	// original authority: JoeBlow:Pass123@Example.COM
	// sanitized authority:                Example.COM
}
