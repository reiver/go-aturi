package aturi

import (
	"strings"
)

// Join returns the AT-URI made up of the provided 'authority', 'collection', 'rkey', 'query', and 'fragment'.
func Join(authority string, collection string, rkey string, query string, fragment string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	// The URI scheme is at, and an authority part preceded with double slashes is always required, so the URI always starts at://
	p = append(p, "at://"...)

	if "" != authority {
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
