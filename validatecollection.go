package aturi

import (
	"github.com/reiver/go-nsid"
)

func ValidateCollection(value string) error {
	return nsid.Validate(value)
}
