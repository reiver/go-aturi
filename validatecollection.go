package aturi

import (
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-nsid"
)

// ValidateCollection returns an error if the AT-URI 'collection' is invalid.
// It returns nil if the AT-URI collection is valid.
//
// NOTE THAT THIS IS NOT VALIDATING AN AT-URI, BUT IS INSTEAD VALIDATING AN AT-URI 'collection'.
//
// The validation rules for an AT-URI are defined here:
// https://atproto.com/specs/at-uri-scheme
//
// An AT-URI 'collection' must be an NSID, and its validation rules are defined at:
// https://atproto.com/specs/nsid
func ValidateCollection(collection string) error {
	return nsid.Validate(collection)
}

func validateCollection(collection string, uri string) error {
	if err := ValidateCollection(collection); nil != err {
		return erorr.Errorf("aturi: AT-URI %q has a collection %q that is not a valid NSID: %w", uri, collection, err)
	}

	return nil
}
