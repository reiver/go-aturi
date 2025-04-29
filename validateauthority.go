package aturi

import (
	"strings"

	"github.com/reiver/go-erorr"
)

const (
	errEmptyAuthority    = erorr.Error("aturi: empty authority")
	errAtSignInAuthority = erorr.Error("aturi: authority may not have an \"@\" in it")
)

// ValidateAuthority returns an error if the AT-URI authority is invalid.
// It returns nil if the AT-URI authority is valid.
//
// The validation rules for an AT-URI are defined here:
// https://atproto.com/specs/at-uri-scheme
func ValidateAuthority(authority string) error {

	if "" == authority {
		return errEmptyAuthority
	}

	{
		const disallowed string = "@"

		if strings.Contains(authority, disallowed) {
			return errAtSignInAuthority
		}
	}

	return nil
}
