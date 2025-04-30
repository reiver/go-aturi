package aturi

import (
	"github.com/reiver/go-erorr"
)

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
