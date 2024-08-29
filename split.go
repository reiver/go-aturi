package aturi

import (
	"strings"

	"github.com/reiver/go-erorr"
)

// Split returns the 'authority', 'collection', 'rkey', 'query', and 'fragment' of at AT-URI.
//
// For example:
//
//	var uri string = "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y"
//	
//	authority, collection, rkey, query, fragment, err := aturl.Split(uri)
//	if nil != err{
//		return err
//	}
//	
//	// authority  == "did:plc:scewmn2pl3oz36mxme2b6czz"
//	// collection == "com.example.foorBar"
//	// rkey       == "3jui7kd54zh2y"
//	// query      == ""
//	// fragment   == ""
func Split(uri string) (authority string, collection string, rkey string, query string, fragment string, err error) {
	if "" == uri {
		return "", "", "", "", "", errEmptyURI
	}

	var str string = uri
	{
		const prefix string = "at://"

		if !strings.HasPrefix(str, prefix) {
			return "", "", "", "", "", erorr.Errorf("aturi: URI %q is not an at-uri because it does not begin with %q", uri, prefix)
		}

		str = str[len(prefix):]
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

		if "" == authority {
			return "", "", "", "", "", erorr.Errorf("aturi: URI %q has an empty 'authority'", uri)
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
