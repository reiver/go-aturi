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


	if err := validate(uri); nil != err {
		return err
	}

	authority, collection, _, _, _, err := Split(uri)
	if nil != err {
		return err
	}

	if err := validateAuthority(authority, uri); nil != err {
		 return err
	}
	if err := validateCollection(collection, uri); nil != err {
		 return err
	}

	return nil
}

func validate(uri string) error {

	{
		const max int = 8192 // == 8 kilobytes == 8 × 1 kilobyte == 8 × 1024 bytes == 8 × 2¹⁰ bytes

		var length int = len(uri)

		if max < length {
			return erorr.Errorf("aturi: URI is %d bytes long but an AT-URI may not be more than %d bytes long", length, max)
		}
	}

	return nil
}
