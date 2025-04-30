package aturi_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-aturi"
)

func TestValidateGenerically(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: "aturi: empty URI",
		},



		{
			URI: "a",
			ExpectedError: `aturi: URI "a" is not an AT-URI because it does not begin with "at:"`,
		},
		{
			URI: "at",
			ExpectedError: `aturi: URI "at" is not an AT-URI because it does not begin with "at:"`,
		},
		{
			URI: "at:",
			ExpectedError: `aturi: AT-URI "at:" is not valid because it does not have "//" after "at:" — too short`,
		},



		{
			URI: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/" + strings.Repeat("0123456789ABCDEFGHIJKLMNOPQRSTUV", 256)[len("at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/")-1:],
			ExpectedError: `aturi: URI is 8193 bytes long but an AT-URI may not be more than 8192 bytes long`,
		},
	}

	for testNumber, test := range tests {

		err := aturi.ValidateGenerically(test.URI)
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

func TestValidateGenerically_ok(t *testing.T) {

	tests := []struct{
		URI string
	}{
		{
			URI: "at://",
		},



		{
			URI: "at://foo.com/com.example.foo/123",
		},



		{
			URI: "at://foo.com/example/123",
		},
		{
			URI: "at://computer",
		},
		{
			URI: "at://example.com:3000",
		},



		{
			URI: "at://",
		},
		{
			URI: "at:///",
		},
		{
			URI: "at://?",
		},
		{
			URI: "at://#",
		},
		{
			URI: "at://?#",
		},



		{
			URI: "at://user:pass@foo.com",
		},
		{
			URI: "at://user:pass@Foo.COM",
		},
		{
			URI: "at://@",
		},
		{
			URI: "at://@example",
		},
		{
			URI: "at://@example.com",
		},



		{
			URI: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/" + strings.Repeat("0123456789ABCDEFGHIJKLMNOPQRSTUV", 256)[len("at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/"):],
		},
	}


	for testNumber, test := range tests {

		err := aturi.ValidateGenerically(test.URI)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %q", test.URI)
			continue
		}

	}
}
