package aturi

import (
	gourl "net/url"
	"strings"

	"github.com/reiver/go-erorr"
)

// Split returns the 'authority', 'collection', 'rkey', 'query', and 'fragment' of at AT-URI.
//
// A 'collection' is an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y"
//	
//	authority, collection, rkey, query, fragment, err := aturi.Split(uri)
//	if nil != err {
//		return err
//	}
//	
//	// authority  == "did:plc:scewmn2pl3oz36mxme2b6czz"
//	// collection == "com.example.foorBar"
//	// rkey       == "3jui7kd54zh2y"
//	// query      == ""
//	// fragment   == ""
//
// Split does NOT normalize the returned values.
// If you are not sure whether to use Split or [SplitAndNormalize], use [SplitAndNormalize].
func Split(uri string) (authority string, collection string, rkey string, query string, fragment string, err error) {
	if "" == uri {
		return "", "", "", "", "", errEmptyURI
	}

	var str string = uri

	{
		const prefix string = "at:"

		var lenprefix int = len(prefix)
		var lenstr int = len(str)
		if lenstr < lenprefix {
			return "", "", "", "", "", erorr.Errorf("aturi: URI %q is not an AT-URI because it does not begin with %q", uri, prefix)
		}

		var beginning string = uri[:lenprefix]

		if strings.ToLower(beginning) != prefix {
			return "", "", "", "", "", erorr.Errorf("aturi: URI %q is not an AT-URI because it does not begin with %q", uri, prefix)
		}

		str = str[lenprefix:]
	}

	{
		const prefix string = "//"

		var lenprefix int = len(prefix)
		var lenstr int = len(str)
		if lenstr < lenprefix {
			return "", "", "", "", "", erorr.Errorf("aturi: AT-URI %q is not valid because it does not have %q after \"at:\" — too short", uri, prefix)
		}

		var beginning string = str[:lenprefix]

		if beginning != prefix {
			return "", "", "", "", "", erorr.Errorf("aturi: AT-URI %q is not valid because it does not have %q after \"at:\"", uri, prefix)
		}

		str = str[lenprefix:]
	}

	// authority
	{
		var index int = strings.Index(str, "/")
		if index < 0 {
			index = strings.Index(str, "?")
			if index < 0 {
				index = strings.Index(str, "#")
			}
		}
		switch {
		case index < 0:
			authority = str
			str = ""
		default:
			authority = str[:index]
			str = str[index:]
		}

		{
			unescaped, err := gourl.QueryUnescape(authority)
			if nil != err {
				return "", "", "", "", "", erorr.Errorf("aturi: problem hex-decoding URI %q: %w", uri, err)
			}
			authority = unescaped
		}
	}

	switch str {
	case "", "/", "?", "#":
		return
	}

	// collection
	{
		const prefix string = "/"

		if strings.HasPrefix(str, prefix)  {
			str = str[len(prefix):]

			var index int = strings.Index(str, "/")
			if index < 0 {
				index = strings.Index(str, "?")
				if index < 0 {
					index = strings.Index(str, "#")
				}
			}

			switch {
			case index < 0:
				collection = str
				str = ""
			default:
				collection = str[:index]
				str = str[index:]
			}
		}
	}

	switch str {
	case "", "/", "?", "#":
		return
	}

	// rkey
	{
		const prefix string = "/"

		if strings.HasPrefix(str, prefix)  {
			str = str[len(prefix):]

			var index int = strings.Index(str, "?")
			if index < 0 {
				index = strings.Index(str, "#")
			}

			switch {
			case index < 0:
				rkey = str
				str = ""
			default:
				rkey = str[:index]
				str = str[index:]
			}
		}

	}

	switch str {
	case "", "?", "#":
		return
	}

	// query
	{
		const prefix string = "?"

		if strings.HasPrefix(str, prefix)  {
			str = str[len(prefix):]

			var index int = strings.Index(str, "#")

			switch {
			case index < 0:
				query = str
				str = ""
			default:
				query = str[:index]
				str = str[index:]
			}
		}
	}

	switch str {
	case "", "#":
		return
	}

	// fragment
	{
		const prefix string = "#"

		if strings.HasPrefix(str, prefix)  {
			str = str[len(prefix):]

			fragment = str
		}
	}

	return
}
