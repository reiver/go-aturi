package aturi_test

import (
	"testing"

	"github.com/reiver/go-aturi"
)

func TestNormalize(t *testing.T) {

	tests := []struct{
		URI string
		Expected string
	}{
		{
			URI:      "a",
			Expected: "a",
		},
		{
			URI:      "A",
			Expected: "A",
		},



		{
			URI:      "http://example.com?wow#",
			Expected: "http://example.com?wow#",
		},
		{
			URI:      "HTTP://Example.COM?wow#",
			Expected: "HTTP://Example.COM?wow#",
		},



		{
			URI:      "at:",
			Expected: "at://",
		},
		{
			URI:      "aT:",
			Expected: "at://",
		},
		{
			URI:      "At:",
			Expected: "at://",
		},
		{
			URI:      "AT:",
			Expected: "at://",
		},



		{
			URI:      "at://",
			Expected: "at://",
		},
		{
			URI:      "aT://",
			Expected: "at://",
		},
		{
			URI:      "At://",
			Expected: "at://",
		},
		{
			URI:      "AT://",
			Expected: "at://",
		},



		{
			URI:      "at:///",
			Expected: "at://",
		},
		{
			URI:      "aT:///",
			Expected: "at://",
		},
		{
			URI:      "At:///",
			Expected: "at://",
		},
		{
			URI:      "AT:///",
			Expected: "at://",
		},



		{
			URI:      "at:////",
			Expected: "at://",
		},
		{
			URI:      "aT:////",
			Expected: "at://",
		},
		{
			URI:      "At:////",
			Expected: "at://",
		},
		{
			URI:      "AT:////",
			Expected: "at://",
		},



		{
			URI:      "at://example.com",
			Expected: "at://example.com",
		},
		{
			URI:      "aT://example.com",
			Expected: "at://example.com",
		},
		{
			URI:      "At://example.com",
			Expected: "at://example.com",
		},
		{
			URI:      "AT://example.com",
			Expected: "at://example.com",
		},



		{
			URI:      "at://Example.COM",
			Expected: "at://example.com",
		},
		{
			URI:      "aT://Example.COM",
			Expected: "at://example.com",
		},
		{
			URI:      "At://Example.COM",
			Expected: "at://example.com",
		},
		{
			URI:      "AT://Example.COM",
			Expected: "at://example.com",
		},



		{
			URI:      "at://did:plc:scewmn2pl3oz36mxme2b6czz",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "aT://did:plc:scewmn2pl3oz36mxme2b6czz",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "At://did:plc:scewmn2pl3oz36mxme2b6czz",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "AT://did:plc:scewmn2pl3oz36mxme2b6czz",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:      "at://Example.COM/",
			Expected: "at://example.com",
		},
		{
			URI:      "aT://Example.COM/",
			Expected: "at://example.com",
		},
		{
			URI:      "At://Example.COM/",
			Expected: "at://example.com",
		},
		{
			URI:      "AT://Example.COM/",
			Expected: "at://example.com",
		},



		{
			URI:      "at://did:plc:scewmn2pl3oz36mxme2b6czz/",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "aT://did:plc:scewmn2pl3oz36mxme2b6czz/",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "At://did:plc:scewmn2pl3oz36mxme2b6czz/",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "AT://did:plc:scewmn2pl3oz36mxme2b6czz/",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:      "at://Example.COM?",
			Expected: "at://example.com",
		},
		{
			URI:      "aT://Example.COM?",
			Expected: "at://example.com",
		},
		{
			URI:      "At://Example.COM?",
			Expected: "at://example.com",
		},
		{
			URI:      "AT://Example.COM?",
			Expected: "at://example.com",
		},



		{
			URI:      "at://did:plc:scewmn2pl3oz36mxme2b6czz?",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "aT://did:plc:scewmn2pl3oz36mxme2b6czz?",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "At://did:plc:scewmn2pl3oz36mxme2b6czz?",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "AT://did:plc:scewmn2pl3oz36mxme2b6czz?",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:      "at://Example.COM#",
			Expected: "at://example.com",
		},
		{
			URI:      "aT://Example.COM#",
			Expected: "at://example.com",
		},
		{
			URI:      "At://Example.COM#",
			Expected: "at://example.com",
		},
		{
			URI:      "AT://Example.COM#",
			Expected: "at://example.com",
		},



		{
			URI:      "at://did:plc:scewmn2pl3oz36mxme2b6czz#",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "aT://did:plc:scewmn2pl3oz36mxme2b6czz#",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "At://did:plc:scewmn2pl3oz36mxme2b6czz#",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			URI:      "AT://did:plc:scewmn2pl3oz36mxme2b6czz#",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:      "at://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar",
		},
		{
			URI:      "aT://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar",
		},
		{
			URI:      "At://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar",
		},
		{
			URI:      "AT://did:plc:scewmn2pl3oz36mxme2b6czz/COM.Example.fooBar",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar",
		},

		{
			URI:      "at://TEST.Example/COM.Example.fooBar",
			Expected: "at://test.example/com.example.fooBar",
		},
		{
			URI:      "aT://TEST.Example/COM.Example.fooBar",
			Expected: "at://test.example/com.example.fooBar",
		},
		{
			URI:      "At://TEST.Example/COM.Example.fooBar",
			Expected: "at://test.example/com.example.fooBar",
		},
		{
			URI:      "AT://TEST.Example/COM.Example.fooBar",
			Expected: "at://test.example/com.example.fooBar",
		},
	}

	for testNumber, test := range tests {

		actual := aturi.Normalize(test.URI)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual normalized AT-URL is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("AT-URI:   %s", test.URI)
			continue
		}
	}
}
