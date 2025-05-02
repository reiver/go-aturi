package aturi

// Normalize returns the normalized form of an AT-URI, as defined in:
// https://atproto.com/specs/at-uri-scheme
//
// Normalize does NOT validate the AT-URI.
// To validate, call [Validate].
//
// An example of a non-normalized AT-URI would be:
//
//	at://VIDEO.archive.ORG/COM.Example.fooBar
//
// Normalizing that non-normalized AT-URI would result in:
//
//	at://video.archive.org/com.example.fooBar
func Normalize(uri string) string {

	if err := ValidateScheme(uri); nil != err {
		return uri
	}

	authority, collection, rkey, query, fragment, err := Split(uri)
	if nil != err {
		return uri
	}

	return Join(authority, collection, rkey, query, fragment)
}
