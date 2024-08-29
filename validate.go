package aturi

// Validate returns an error if the AT-URI is invalid.
// It returns nil if the AT-URI is valid.
func Validate(uri string) error {
	_, _, _, _, _, err := Split(uri)
	return err
}
