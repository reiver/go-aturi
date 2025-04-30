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
// You would use ValidateScheme if you wanted to be very liberal in what you accept as a valid AT-URI, including not caring if the AT-URI has an 'authority' or not.
// I.e., this is the minimum amount of validation you can do to validate an AT-URI.
//
// Note that you are passing ValidateScheme the whole URI, and not just the scheme.
//
// For a more thorough validation of the whole AT-URI instead use [Validate].
//
// Alternatively, to validate a bit more than ValidateScheme, without being as thorough as [Validate], instead use [ValidatePrefix].
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
