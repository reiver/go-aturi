package aturi_test

import (
	"testing"

	"github.com/reiver/go-aturi"
)

func TestNormalizeAuthority(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
		},



		{
			Value:    "example.com",
			Expected: "example.com",
		},



		{
			Value:    "example.coM",
			Expected: "example.com",
		},
		{
			Value:    "example.cOm",
			Expected: "example.com",
		},
		{
			Value:    "example.Com",
			Expected: "example.com",
		},
		{
			Value:    "examplE.com",
			Expected: "example.com",
		},
		{
			Value:    "exampLe.com",
			Expected: "example.com",
		},
		{
			Value:    "examPle.com",
			Expected: "example.com",
		},
		{
			Value:    "exaMple.com",
			Expected: "example.com",
		},
		{
			Value:    "exAmple.com",
			Expected: "example.com",
		},
		{
			Value:    "eXample.com",
			Expected: "example.com",
		},
		{
			Value:    "Example.com",
			Expected: "example.com",
		},



		{
			Value:    "Example.COM",
			Expected: "example.com",
		},



		{
			Value:    "apple.BANANA.Cherry",
			Expected: "apple.banana.cherry",
		},



		{
			Value:    "ABC😈123",
			Expected: "abc😈123",
		},



		{
			Value:    "JoeBlow:pass123@Example.COM",
			Expected: "JoeBlow:pass123@example.com",
		},
		{
			Value:    "@Example.COM",
			Expected: "@example.com",
		},
		{
			Value:    "JoeBlow:pass123@",
			Expected: "JoeBlow:pass123@",
		},
	}

	for testNumber, test := range tests {

		actual := aturi.NormalizeAuthority(test.Value)

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
