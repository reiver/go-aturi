package aturi

import (
	"strings"

	"golang.org/x/net/idna"
)

// Join returns the AT-URI made up of the provided 'authority', 'collection', 'rkey', 'query', and 'fragment'.
//
// I.e., you use Join to create an AT-URI.
//
// Join will take care of the various complexities of properly creating a (normalized) AT-URI.
//
// (Most people will set 'query' and 'fragment' to the empty string.)
//
// Join normalizes the 'authority' and 'collection' before joining — by internally calling [NormalizeAuthority] and [NormalizeCollection] respectively.
// So, you do not need to call [NormalizeAuthority] and [NormalizeCollection] yourself, before passing the 'authority' and 'collection' to Join.
//
// Join also sanitizes 'authority' before joining — by internally calling [SanitizeAuthority].
func Join(authority string, collection string, rkey string, query string, fragment string) string {
	authority = SanitizeAuthority(authority)

	authority  = NormalizeAuthority(authority)
	collection = NormalizeCollection(collection)

	return join(authority, collection, rkey, query, fragment)
}

func join(authority string, collection string, rkey string, query string, fragment string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	// The URI scheme is at, and an authority part preceded with double slashes is always required, so the URI always starts at://
	p = append(p, "at://"...)

	if "" != authority {
		{
			punycode, err := idna.ToASCII(authority)
			if nil == err {
				authority = punycode
			}
		}

		p = append(p, encodeAtSign(encodeSolidus(encodeQuestionMark(encodeNumberSign(encodePercentSign(authority)))))...)
	}

	if  "" != collection || "" != rkey {
		p = append(p, "/"...)

		if "" != collection {
			p = append(p, encodeSolidus(encodeQuestionMark(encodeNumberSign(encodePercentSign(collection))))...)
		}

		if "" != rkey {
			p = append(p, "/"...)
			p = append(p, encodeQuestionMark(encodeNumberSign(encodePercentSign(rkey)))...)
		}
	}

	if "" != query {
		p = append(p, "?"...)
		p = append(p, encodeNumberSign(encodePercentSign(query))...)
	}

	if "" != fragment {
		p = append(p, "#"...)
		p = append(p, fragment...)
	}

	return string(p)
}

func encodeAtSign(str string) string {
	return strings.ReplaceAll(str,  "@", "%40")
}

func encodeNumberSign(str string) string {
	return strings.ReplaceAll(str,  "#", "%23")
}

func encodePercentSign(str string) string {
	return strings.ReplaceAll(str,  "%", "%25")
}

func encodeQuestionMark(str string) string {
	return strings.ReplaceAll(str,  "?", "%3F")
}

func encodeSolidus(str string) string {
	return strings.ReplaceAll(str,  "/", "%2F")
}
