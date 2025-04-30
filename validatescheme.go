package aturi

import (
	"strings"

	"github.com/reiver/go-erorr"
)

const schemePrefix string = "at:"
const lenSchemePrefix int = len(schemePrefix)

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
		var lenuri int = len(uri)
		if lenuri < lenSchemePrefix {
			return erorr.Errorf("aturi: URI %q is not an AT-URI because it does not begin with %q", uri, schemePrefix)
		}

                var beginning string = uri[:lenSchemePrefix]

                if strings.ToLower(beginning) != schemePrefix {
			return erorr.Errorf("aturi: URI %q is not an AT-URI because it does not begin with %q", uri, schemePrefix)
                }
	}

	return nil
}
