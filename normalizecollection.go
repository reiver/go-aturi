package aturi

import (
	"github.com/reiver/go-nsid"
)

// NormalizeCollection returns the normalized form of an AT-URI collection, as defined in:
// https://atproto.com/specs/at-uri-scheme
//
// An AT-URI collection is an NSID, as defined at:
// https://atproto.com/specs/nsid
//
// The AT-URI collection (such as "com.example.fooBar") is part of an AT-URI (such as "at://archive.org/com.example.fooBar/123").
//
// An example of a non-normalized AT-URI collection would be "COM.Example.fooBar".
// Normalizing that non-normalized AT-URI collection would result in "com.example.fooBar".
//
// Note that if you want to normalize a whole AT-URI rather than just a collection, then instead use [Normalize].
func NormalizeCollection(value string) string {
	return nsid.Normalize(value)
}
