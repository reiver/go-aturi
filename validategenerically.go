package aturi

import (
	"github.com/reiver/go-erorr"
)

// ValidateGenerically only validates an AT-URI very generically.
//
// So, it checks to see if the URI starts with "at://", and make sure that the AT-URI is not more than 8192 bytes long (i.e., 8 kilobytes long).
// And, that is it.
//
// You would use ValidateGenerically if you wanted to be very liberal in what you accept as a valid AT-URI, while still enforcing the length limit.
//
// ValidateGenerically calls [ValidatePrefix] internally.
//
// For a more thorough validation of the whole AT-URI instead use [Validate].
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
