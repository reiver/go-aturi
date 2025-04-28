package aturi

import (
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-nsid"
)

// Validate returns an error if the AT-URI is invalid.
// It returns nil if the AT-URI is valid.
//
// The validation rules for an AT-URI are defined here:
// https://atproto.com/specs/at-uri-scheme
func Validate(uri string) error {
	_, collection, _, _, _, err := Split(uri)
	if nil != err {
		return err
	}

	if 0 < len(collection) {
		if err := nsid.Validate(collection); nil != err {
			return erorr.Errorf("aturi: AT-URI %q has a collection %q that is not a valid NSID: %w", uri, collection, err)
		}
	}

	return nil
}
