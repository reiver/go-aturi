package aturi

import (
	"github.com/reiver/go-erorr"
)

// ValidateGenerically only validates an AT-URI in a very general form.
//
// So, it checks to see if the URI starts with "at://", and make sure that the AT-URI is not more than 8192 bytes long (i.e., 8 kilobytes long).
// And, that is it.
//
// ValidateGenerically calls [ValidatePrefix] internally.
//
// For more thorough validation of the whole AT-URI instead use [Validate].
func ValidateGenerically(uri string) error {

	if err := ValidatePrefix(uri); nil != err {
		return err
	}

	{
		const max int = 8192 // == 8 kilobytes == 8 × 1 kilobyte == 8 × 1024 bytes == 8 × 2¹⁰ bytes

		var length int = len(uri)

		if max < length {
			return erorr.Errorf("aturi: URI is %d bytes long but an AT-URI may not be more than %d bytes long", length, max)
		}
	}

	return nil
}
