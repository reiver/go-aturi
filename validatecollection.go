package aturi

import (
	"github.com/reiver/go-nsid"
)

// ValidateCollection returns an error if the AT-URI collection is invalid.
// It returns nil if the AT-URI collection is valid.
//
// The validation rules for an AT-URI are defined here:
// https://atproto.com/specs/at-uri-scheme
//
// An AT-URI collection must be an NSID, and follow its validation rules, as defined at:
// https://atproto.com/specs/nsid
func ValidateCollection(value string) error {
	return nsid.Validate(value)
}
