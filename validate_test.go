package aturi_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-aturi"
)

func TestValidate(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: "aturi: empty URI",
		},



                {
                        URI: "at://foo.com/example/123",
                        ExpectedError: `aturi: AT-URI "at://foo.com/example/123" has a collection "example" that is not a valid NSID: nsid: nsid ("example") should have at least 3 segments but actually has 1`,
                },



		{
			URI: "at://",
			ExpectedError: `aturi: AT-URI "at://" has an empty 'authority'`,
		},
		{
			URI: "at:///",
			ExpectedError: `aturi: AT-URI "at:///" has an empty 'authority'`,
		},
		{
			URI: "at://?",
			ExpectedError: `aturi: AT-URI "at://?" has an empty 'authority'`,
		},
		{
			URI: "at://#",
			ExpectedError: `aturi: AT-URI "at://#" has an empty 'authority'`,
		},
		{
			URI: "at://?#",
			ExpectedError: `aturi: AT-URI "at://?#" has an empty 'authority'`,
		},



		{
			URI: "at://user:pass@foo.com",
			ExpectedError: `aturi: AT-URI "at://user:pass@foo.com" may not have an "@" in its authority "user:pass@foo.com"`,
		},
		{
			URI: "at://user:pass@Foo.COM",
			ExpectedError: `aturi: AT-URI "at://user:pass@Foo.COM" may not have an "@" in its authority "user:pass@Foo.COM"`,
		},
		{
			URI: "at://@",
			ExpectedError: `aturi: AT-URI "at://@" may not have an "@" in its authority "@"`,
		},
		{
			URI: "at://@example",
			ExpectedError: `aturi: AT-URI "at://@example" may not have an "@" in its authority "@example"`,
		},
		{
			URI: "at://@example.com",
			ExpectedError: `aturi: AT-URI "at://@example.com" may not have an "@" in its authority "@example.com"`,
		},


		{
			URI: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/" + strings.Repeat("0123456789ABCDEFGHIJKLMNOPQRSTUV", 256)[len("at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/")-1:],
			ExpectedError: `aturi: URI is 8193 bytes long but an AT-URI may not be more than 8192 bytes long`,
		},
	}

	for testNumber, test := range tests {

		err := aturi.Validate(test.URI)
		if nil == err {
			t.Errorf("For test #%d, expected an error but didn't get one.", testNumber)
			t.Logf("URI: %q", test.URI)
			continue
		}

		expected := test.ExpectedError
		actual := err.Error()

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("URI: %q", test.URI)
			continue
		}
	}
}
