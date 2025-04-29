package aturi

import (
	"strings"
)

// NormalizeAuthority returns the normalized form of an AT-URI authority, as defined in:
// https://atproto.com/specs/at-uri-scheme
//
// The AT-URI authority (such as "example.com") is part of an AT-URI (such as "at://example.com").
//
// In simple language, you can think of an AT-URI authority as being either an Internet domain-name (ex: "example.com") or a DID (ex: "did:plc:scewmn2pl3oz36mxme2b6czz").
//
// An example of a non-normalized AT-URI authority would be "Example.COM".
// Normalizing that non-normalized AT-URI authority would result in "example.com".
//
// Note that if you want to normalize a whole AT-URI rather than just an authority, then instead use [Normalize].
func NormalizeAuthority(value string) string {
	if strings.HasPrefix(value, "did:") {
		return value
	}

	var str string = value

	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
		var index int = strings.Index(str, "@")
		if 0 <= index {
			var userinfo string = str[:index]
			p = append(p, userinfo...)
			p = append(p, '@')

			str = str[index+1:]
		}
	}

	var length int = len(str)

	for i:=0; i<length; i++ {
		var b byte = str[i]

		switch {
		case 'A' <= b && b <= 'Z':
			p = append(p, b-'A'+'a')
		default:
			p = append(p, b)
		}
	}

	return string(p)
}
