package aturi

import (
	"github.com/reiver/go-erorr"
)

// Validate returns an error if the AT-URI is invalid.
// It returns nil if the AT-URI is valid.
//
// The validation rules for an AT-URI are defined here:
// https://atproto.com/specs/at-uri-scheme
func Validate(uri string) error {

	{
		const max int = 8192 // == 8 kilobytes == 8 × 1 kilobyte == 8 × 1024 bytes == 8 × 2¹⁰ bytes

		var length int = len(uri)

		if max < length {
			return erorr.Errorf("aturi: URI is %d bytes long but an AT-URI may not be more than %d bytes long", length, max)
		}
	}

	authority, collection, _, _, _, err := Split(uri)
	if nil != err {
		return err
	}

	if err := ValidateAuthority(authority); nil != err {
		switch {
		case erorr.Is(err, errEmptyAuthority):
			return erorr.Errorf("aturi: AT-URI %q has an empty 'authority'", uri)
		case erorr.Is(err, errAtSignInAuthority):
			return erorr.Errorf("aturi: AT-URI %q may not have an \"@\" in its authority %q", uri, authority)
		default:
			return erorr.Errorf("aturi: AT-URI %q has a authority %q that is not a valid: %w", uri, authority, err)
		}
	}

	if err := ValidateCollection(collection); nil != err {
		 return erorr.Errorf("aturi: AT-URI %q has a collection %q that is not a valid NSID: %w", uri, collection, err)
	}

	return nil
}
