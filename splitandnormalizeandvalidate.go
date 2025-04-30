package aturi

// SplitAndNormalizeAndValidate returns the 'authority', 'collection', 'rkey', 'query', and 'fragment' of at AT-URI.
//
// A 'collection' is an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y"
//	
//	authority, collection, rkey, query, fragment, err := aturi.SplitAndNormalizeAndValidate(uri)
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
// SplitAndNormalize normalizes the returned values.
// SplitAndNormalizeAndValidate also validates, and returns an error if it fails.
//
// If you are not sure whether to use [Split] or [SpliAndNormalize] or SplitAndNormalizeAndValidate, use SplitAndNormalizeAndValidate.
func SplitAndNormalizeAndValidate(uri string) (authority string, collection string, rkey string, query string, fragment string, err error) {
	if err := ValidateGenerically(uri); nil != err {
		return "", "", "", "", "", err
	}

	authority, collection, rkey, query, fragment, err = SplitAndNormalize(uri)
	if nil != err {
		return authority, collection, rkey, query, fragment, err
	}

	if err := validateAuthority(authority, uri); nil != err {
		return authority, collection, rkey, query, fragment, err
	}
	if "" != collection {
		if err := validateCollection(collection, uri); nil != err {
			return authority, collection, rkey, query, fragment, err
		}
	}

	return authority, collection, rkey, query, fragment, err
}
