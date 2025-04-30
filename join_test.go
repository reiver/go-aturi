package aturi_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-aturi"
)

func TestJoin(t *testing.T) {

	tests := []struct{
		Authority string
		Collection string
		RKey string
		Query string
		Fragment string
		Expected string
	}{
		{
			Expected: "at://",
		},



		{
			Collection:     "com.example.foo",
			Expected: "at:///com.example.foo",
		},
		{
			RKey:            "123",
			Expected: "at:////123",
		},
		{
			Query:          "a=bb&ccc=dddd",
			Expected: "at://?a=bb&ccc=dddd",
		},
		{
			Fragment:       "some",
			Expected: "at://#some",
		},



		{
			Authority:     "foo.com",
			Collection:            "com.example.foo",
			RKey:                                  "123",
			Expected: "at://foo.com/com.example.foo/123",
		},
		{
			Authority:     "foo.com",
			Collection:            "com.example.foo",
			Expected: "at://foo.com/com.example.foo",
		},
		{
			Authority:     "foo.com",
			RKey:                   "123",
			Expected: "at://foo.com//123",
		},
		{
			RKey:            "123",
			Expected: "at:////123",
		},



		{
			Authority:     "localhost",
			Expected: "at://localhost",
		},
		{
			Authority:     "example.com",
			Expected: "at://example.com",
		},
		{
			Authority:     "example.com.",
			Expected: "at://example.com.",
		},
		{
			Authority:     "apple.banana.cherry",
			Expected: "at://apple.banana.cherry",
		},
		{
			Authority:     "xn--ugbaf6g.example",
			Expected: "at://xn--ugbaf6g.example",
		},
		{
			Authority:     "did:plc:scewmn2pl3oz36mxme2b6czz",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz",
		},
		{
			Authority:     "1/2",
			Expected: "at://1%2F2",
		},
		{
			Authority:     "a?b",
			Expected: "at://a%3Fb",
		},
		{
			Authority:     "x#y",
			Expected: "at://x%23y",
		},
		{
			Authority:     "j%k",
			Expected: "at://j%25k",
		},



		{
			Authority:     "localhost",
			Collection:              "com.example.foorBar",
			Expected: "at://localhost/com.example.foorBar",
		},
		{
			Authority:     "example.com",
			Collection:                "com.example.foorBar",
			Expected: "at://example.com/com.example.foorBar",
		},
		{
			Authority:     "example.com.",
			Collection:                 "com.example.foorBar",
			Expected: "at://example.com./com.example.foorBar",
		},
		{
			Authority:     "apple.banana.cherry",
			Collection:                        "com.example.foorBar",
			Expected: "at://apple.banana.cherry/com.example.foorBar",
		},
		{
			Authority:     "xn--ugbaf6g.example",
			Collection:                        "com.example.foorBar",
			Expected: "at://xn--ugbaf6g.example/com.example.foorBar",
		},
		{
			Authority:     "did:plc:scewmn2pl3oz36mxme2b6czz",
			Collection:                                     "com.example.foorBar",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar",
		},
		{
			Authority:     "1/2",
			Collection:          "com.example.foorBar",
			Expected: "at://1%2F2/com.example.foorBar",
		},
		{
			Collection:          "com.example.foorBar",
			Authority:     "a?b",
			Expected: "at://a%3Fb/com.example.foorBar",
		},
		{
			Authority:     "x#y",
			Collection:          "com.example.foorBar",
			Expected: "at://x%23y/com.example.foorBar",
		},
		{
			Authority:     "j%k",
			Collection:          "com.example.foorBar",
			Expected: "at://j%25k/com.example.foorBar",
		},




		{
			Authority:     "localhost",
			Collection:              "com.example.foorBar",
			RKey:                                        "3jui7kd54zh2y",
			Expected: "at://localhost/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "example.com",
			Collection:                "com.example.foorBar",
			RKey:                                          "3jui7kd54zh2y",
			Expected: "at://example.com/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "example.com.",
			Collection:                 "com.example.foorBar",
			RKey:                                           "3jui7kd54zh2y",
			Expected: "at://example.com./com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "apple.banana.cherry",
			Collection:                        "com.example.foorBar",
			RKey:                                                  "3jui7kd54zh2y",
			Expected: "at://apple.banana.cherry/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "xn--ugbaf6g.example",
			Collection:                        "com.example.foorBar",
			RKey:                                                  "3jui7kd54zh2y",
			Expected: "at://xn--ugbaf6g.example/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "did:plc:scewmn2pl3oz36mxme2b6czz",
			Collection:                                     "com.example.foorBar",
			RKey:                                                               "3jui7kd54zh2y",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "1/2",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Expected: "at://1%2F2/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Collection:          "com.example.foorBar",
			Authority:     "a?b",
			RKey:                                    "3jui7kd54zh2y",
			Expected: "at://a%3Fb/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "x#y",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Expected: "at://x%23y/com.example.foorBar/3jui7kd54zh2y",
		},
		{
			Authority:     "j%k",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Expected: "at://j%25k/com.example.foorBar/3jui7kd54zh2y",
		},



		{
			Authority:     "localhost",
			Collection:              "com.example.foorBar",
			RKey:                                        "3jui7kd54zh2y",
			Query:                                                     "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://localhost/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "example.com",
			Collection:            "com.example.foorBar",
			RKey:                                          "3jui7kd54zh2y",
			Query:                                                       "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://example.com/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "example.com.",
			Collection:             "com.example.foorBar",
			RKey:                                           "3jui7kd54zh2y",
			Query:                                                        "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://example.com./com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "apple.banana.cherry",
			Collection:                    "com.example.foorBar",
			RKey:                                                  "3jui7kd54zh2y",
			Query:                                                               "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://apple.banana.cherry/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "xn--ugbaf6g.example",
			Collection:                    "com.example.foorBar",
			RKey:                                                  "3jui7kd54zh2y",
			Query:                                                               "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://xn--ugbaf6g.example/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "did:plc:scewmn2pl3oz36mxme2b6czz",
			Collection:                                 "com.example.foorBar",
			RKey:                                                               "3jui7kd54zh2y",
			Query:                                                                            "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "1/2",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://1%2F2/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Collection:          "com.example.foorBar",
			Authority:     "a?b",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://a%3Fb/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "x#y",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://x%23y/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},
		{
			Authority:     "j%k",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Expected: "at://j%25k/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
		},




		{
			Expected: "at://localhost/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			Authority:     "localhost",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                               "path(/apple/banana/cherry)",
		},
		{
			Expected: "at://example.com/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			Authority:     "example.com",
			Collection:            "com.example.foorBar",
			RKey:                                      "3jui7kd54zh2y",
			Query:                                                   "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                                 "path(/apple/banana/cherry)",
		},
		{
			Expected: "at://example.com./com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			Authority:     "example.com.",
			Collection:             "com.example.foorBar",
			RKey:                                       "3jui7kd54zh2y",
			Query:                                                    "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                                  "path(/apple/banana/cherry)",
		},
		{
			Expected: "at://apple.banana.cherry/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			Authority:     "apple.banana.cherry",
			Collection:                    "com.example.foorBar",
			RKey:                                              "3jui7kd54zh2y",
			Query:                                                           "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                                         "path(/apple/banana/cherry)",
		},
		{
			Expected: "at://xn--ugbaf6g.example/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			Authority:     "xn--ugbaf6g.example",
			Collection:                    "com.example.foorBar",
			RKey:                                              "3jui7kd54zh2y",
			Query:                                                           "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                                         "path(/apple/banana/cherry)",
		},
		{
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			Authority:     "did:plc:scewmn2pl3oz36mxme2b6czz",
			Collection:                                 "com.example.foorBar",
			RKey:                                                           "3jui7kd54zh2y",
			Query:                                                                        "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                                                      "path(/apple/banana/cherry)",
		},
		{
			Authority:     "1/2",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                               "path(/apple/banana/cherry)",
			Expected: "at://1%2F2/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
		},
		{
			Collection:          "com.example.foorBar",
			Authority:     "a?b",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                               "path(/apple/banana/cherry)",
			Expected: "at://a%3Fb/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
		},
		{
			Authority:     "x#y",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                               "path(/apple/banana/cherry)",
			Expected: "at://x%23y/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
		},
		{
			Authority:     "j%k",
			Collection:          "com.example.foorBar",
			RKey:                                    "3jui7kd54zh2y",
			Query:                                                 "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                               "path(/apple/banana/cherry)",
			Expected: "at://j%25k/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
		},



		{
			Authority:     "did:plc:scewmn2pl3oz36mxme2b6czz",
			Collection:                                 "com.example.foorBar",
			RKey:                                                           "QRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV",
			Expected: "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/" + strings.Repeat("0123456789ABCDEFGHIJKLMNOPQRSTUV", 256)[len("at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.foorBar/"):],
		},









		{
			Authority:     "Archive.ORG",
			Collection:                "COM.Example.foorBar",
			RKey:                                          "3jui7kd54zh2y",
			Query:                                                       "once=1&twice=2&thrice=3&fource=4",
			Fragment:                                                                                     "path(/apple/banana/cherry)",
			Expected: "at://archive.org/com.example.foorBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
		},









		{
			Authority: "JoeBlow:Pass123@Archive.ORG",
			Expected:             "at://archive.org",
		},



		{
			Authority: "JoeBlow:Pass123@😈.ORG",
			Expected:             "at://xn--m28h.org",
		},
	}

	for testNumber, test := range tests {

		actual := aturi.Join(test.Authority, test.Collection, test.RKey, test.Query, test.Fragment)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual 'AT-URI' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("AUTHORITY:  %q", test.Authority)
				t.Logf("COLLECTION: %q", test.Collection)
				t.Logf("RKEY:       %q", test.RKey)
				t.Logf("QUERY:      %q", test.Query)
				t.Logf("FRAGMENT:   %q", test.Fragment)
				continue
			}
		}
	}
}

