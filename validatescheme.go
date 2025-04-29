package aturi

import (
	"strings"

	"github.com/reiver/go-erorr"
)

// ValidateScheme only validates the scheme of an URI.
//
// So, it checks to see if the URI starts with "at:".
// And, that is it.
//
// For more thorough validation of the whole AT-URI instead use [Validate].
func ValidateScheme(uri string) error {

	if "" == uri {
		return errEmptyURI
	}

	{
		const prefix string = "at:"
		const lenprefix int = len(prefix)

		var lenuri int = len(uri)
		if lenuri < lenprefix {
			return erorr.Errorf("aturi: URI %q is not an AT-URI because it does not begin with %q", uri, prefix)
		}

                var beginning string = uri[:lenprefix]

                if strings.ToLower(beginning) != prefix {
			return erorr.Errorf("aturi: URI %q is not an AT-URI because it does not begin with %q", uri, prefix)
                }
	}

	return nil
}
