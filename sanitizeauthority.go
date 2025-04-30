package aturi

import (
	"strings"

	"github.com/reiver/go-did"
)

// SanitizeAuthority sanitizes an 'authority'.
//
// (An 'authority' is part of an AT-URI.)
//
// Note that SanitizeAuthority does NOT normalize the 'authority'.
// To normalize the 'authority' use [NormalizeAuthority].
func SanitizeAuthority(authority string) string {
	if nil == did.ValidateScheme(authority) {
		return authority
	}

	{
		var index int = strings.Index(authority, "@")
		if 0 <= index {
                        return authority[index+1:]
		}
	}

	return authority
}


