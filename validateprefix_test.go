package aturi_test

import (
	"testing"

	"github.com/reiver/go-aturi"
)

func TestValidatePrefix(t *testing.T) {

	tests := []struct{
		URI string
	}{
		{
			URI: "at://",
		},
		{
			URI: "aT://",
		},
		{
			URI: "At://",
		},
		{
			URI: "AT://",
		},



		{
			URI: "at://VIDEO.Archive.ORG",
		},
		{
			URI: "aT://VIDEO.Archive.ORG",
		},
		{
			URI: "At://VIDEO.Archive.ORG",
		},
		{
			URI: "AT://VIDEO.Archive.ORG",
		},



		{
			URI: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "aT://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "At://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI: "AT://did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI: "at://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "aT://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "At://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "AT://VIDEO.Archive.ORG/COM.Example.fooBar",
		},



		{
			URI: "at://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
		},
		{
			URI: "aT://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
		},
		{
			URI: "At://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
		},
		{
			URI: "AT://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
		},



		{
			URI: "at://VIDEO.Archive.ORG/COM.Example.fooBar/123",
		},
		{
			URI: "aT://VIDEO.Archive.ORG/COM.Example.fooBar/123",
		},
		{
			URI: "At://VIDEO.Archive.ORG/COM.Example.fooBar/123",
		},
		{
			URI: "AT://VIDEO.Archive.ORG/COM.Example.fooBar/123",
		},



		{
			URI: "at://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar/123",
		},
		{
			URI: "aT://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar/123",
		},
		{
			URI: "At://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar/123",
		},
		{
			URI: "AT://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar/123",
		},



		{
			URI: "AT://user:pass@SOMETHING.Example/banana/wxyz",
		},
	}

	for testNumber, test := range tests {

		err := aturi.ValidatePrefix(test.URI)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}

func TestValidatePrefix_fail(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: `aturi: empty URI`,
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
			URI: "aT",
			ExpectedError: `aturi: URI "aT" is not an AT-URI because it does not begin with "at:"`,
		},
		{
			URI: "At",
			ExpectedError: `aturi: URI "At" is not an AT-URI because it does not begin with "at:"`,
		},
		{
			URI: "AT",
			ExpectedError: `aturi: URI "AT" is not an AT-URI because it does not begin with "at:"`,
		},



		{
			URI: "at:",
			ExpectedError: `aturi: AT-URI "at:" is not valid because it does not have "//" after "at:" — too short`,
		},
		{
			URI: "aT:",
			ExpectedError: `aturi: AT-URI "aT:" is not valid because it does not have "//" after "at:" — too short`,
		},
		{
			URI: "At:",
			ExpectedError: `aturi: AT-URI "At:" is not valid because it does not have "//" after "at:" — too short`,
		},
		{
			URI: "AT:",
			ExpectedError: `aturi: AT-URI "AT:" is not valid because it does not have "//" after "at:" — too short`,
		},



		{
			URI: "http://example.com",
			ExpectedError: `aturi: URI "http://example.com" is not an AT-URI because it does not begin with "at:"`,
		},
	}

	for testNumber, test := range tests {

		err := aturi.ValidatePrefix(test.URI)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("URI: %s", test.URI)
			continue
		}

		actual := err.Error()
		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED-ERROR: %s", expected)
			t.Logf("ACTUAL-ERROR:   %s", actual)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}
