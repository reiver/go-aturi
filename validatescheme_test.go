package aturi_test

import (
	"testing"

	"github.com/reiver/go-aturi"
)

func TestValidateScheme(t *testing.T) {

	tests := []struct{
		URI string
	}{
		{
			URI: "at:",
		},
		{
			URI: "aT:",
		},
		{
			URI: "At:",
		},
		{
			URI: "AT:",
		},



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

		err := aturi.ValidateScheme(test.URI)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}
