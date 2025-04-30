package aturi

// Validate returns an error if the AT-URI is invalid.
// It returns nil if the AT-URI is valid.
//
// The validation rules for an AT-URI are defined here:
// https://atproto.com/specs/at-uri-scheme
func Validate(uri string) error {

	if err := ValidateGenerically(uri); nil != err {
		return err
	}

	authority, collection, _, _, _, err := Split(uri)
	if nil != err {
		return err
	}

	if err := validateAuthority(authority, uri); nil != err {
		 return err
	}
	if "" != collection {
		if err := validateCollection(collection, uri); nil != err {
			 return err
		}
	}

	return nil
}
