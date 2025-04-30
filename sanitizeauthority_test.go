package aturi_test

import (
	"testing"

	"github.com/reiver/go-aturi"
)

func TestSanitizeAuthority(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
		},



		{
			Value:    "did:plc:scewmn2pl3oz36mxme2b6czz",
			Expected: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			Value:    "example.com",
			Expected: "example.com",
		},



		{
			Value:    "example.coM",
			Expected: "example.coM",
		},
		{
			Value:    "example.cOm",
			Expected: "example.cOm",
		},
		{
			Value:    "example.Com",
			Expected: "example.Com",
		},
		{
			Value:    "examplE.com",
			Expected: "examplE.com",
		},
		{
			Value:    "exampLe.com",
			Expected: "exampLe.com",
		},
		{
			Value:    "examPle.com",
			Expected: "examPle.com",
		},
		{
			Value:    "exaMple.com",
			Expected: "exaMple.com",
		},
		{
			Value:    "exAmple.com",
			Expected: "exAmple.com",
		},
		{
			Value:    "eXample.com",
			Expected: "eXample.com",
		},
		{
			Value:    "Example.com",
			Expected: "Example.com",
		},



		{
			Value:    "Example.COM",
			Expected: "Example.COM",
		},



		{
			Value:    "apple.BANANA.Cherry",
			Expected: "apple.BANANA.Cherry",
		},



		{
			Value:    "ABC😈123",
			Expected: "ABC😈123",
		},



		{
			Value:    "JoeBlow:pass123@Example.COM",
			Expected:                 "Example.COM",
		},
		{
			Value:    "@Example.COM",
			Expected:  "Example.COM",
		},
		{
			Value:    "JoeBlow:pass123@",
			Expected: "",
		},
	}

	for testNumber, test := range tests {

		actual := aturi.SanitizeAuthority(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual normalized-authority is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE:    %q", test.Value)
			continue
		}
	}
}
