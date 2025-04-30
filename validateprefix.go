package aturi

import (
	"github.com/reiver/go-erorr"
)

// ValidatePrefix only validates the prefix (i.e., "at://") of an AT-URI.
//
// So, it checks to see if the URI starts with "at://".
// And, that is it.
//
// You would use ValidatePrefix if you wanted to be very liberal in what you accept as a valid AT-URI.
//
// ValidatePrefix calls [ValidateScheme] internally.
//
// For a more thorough validation of the whole AT-URI instead use [Validate].
func ValidatePrefix(uri string) error {

	if err := ValidateScheme(uri); nil != err {
		return err
	}

	str := uri[lenSchemePrefix:]

	{
		const slashSlashPrefix string = "//"
		const lenSlashSlashPrefix int = len(slashSlashPrefix)

		var lenstr int = len(str)
		if lenstr < lenSlashSlashPrefix {
			return erorr.Errorf("aturi: AT-URI %q is not valid because it does not have %q after \"at:\" — too short", uri, slashSlashPrefix)
		}

		var beginning string = str[:lenSlashSlashPrefix]

		if beginning != slashSlashPrefix {
			return erorr.Errorf("aturi: AT-URI %q is not valid because it does not have %q after \"at:\"", uri, slashSlashPrefix)
		}
	}

	return nil
}
