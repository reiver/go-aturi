package aturi_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-aturi"
)

func TestSplitAndNormalizeAndValidate(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedAuthority string
		ExpectedCollection string
		ExpectedRKey string
		ExpectedQuery string
		ExpectedFragment string
	}{
		{
			URI:          "at://foo.com/com.example.foo/123",
			ExpectedAuthority: "foo.com",
			ExpectedCollection:        "com.example.foo",
			ExpectedRKey:                              "123",
		},
		{
			URI:          "at://FOO.Com/COM.Example.foo/123",
			ExpectedAuthority: "foo.com",
			ExpectedCollection:        "com.example.foo",
			ExpectedRKey:                              "123",
		},



		{
                        URI:          "at://1%2F2",
                        ExpectedAuthority: "1/2",
		},
		{
                        URI:          "at://a%3Fb",
                        ExpectedAuthority: "a?b",
		},
		{
                        URI:          "at://x%23y",
                        ExpectedAuthority: "x#y",
		},
		{
                        URI:          "at://j%25k",
                        ExpectedAuthority: "j%k",
		},



		{
			URI:          "AT://localhost",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "AT://example.com",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "AT://example.com.",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "AT://apple.banana.cherry",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "AT://Hello.WORLD",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "AT://xn--ugbaf6g.example",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "AT://did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "At://localhost",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "At://example.com",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "At://example.com.",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "At://apple.banana.cherry",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "At://Hello.WORLD",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "At://xn--ugbaf6g.example",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "At://did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "aT://localhost",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "aT://example.com",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "aT://example.com.",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "aT://apple.banana.cherry",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "aT://Hello.WORLD",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "aT://xn--ugbaf6g.example",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "aT://did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com.",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost/",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com/",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com./",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry/",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD/",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example/",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost?",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com?",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com.?",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry?",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD?",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example?",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz?",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost#",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com#",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com.#",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry#",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD#",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example#",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost/?",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com/?",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com./?",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry/?",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD/?",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example/?",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/?",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost/#",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com/#",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com./#",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry/#",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD/#",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example/#",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost?#",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com?#",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com.?#",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry?#",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD?#",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example?#",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz?#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost/?#",
			ExpectedAuthority: "localhost",
		},
		{
			URI:          "at://example.com/?#",
			ExpectedAuthority: "example.com",
		},
		{
			URI:          "at://example.com./?#",
			ExpectedAuthority: "example.com.",
		},
		{
			URI:          "at://apple.banana.cherry/?#",
			ExpectedAuthority: "apple.banana.cherry",
		},
		{
			URI:          "at://Hello.WORLD/?#",
			ExpectedAuthority: "hello.world",
		},
		{
			URI:          "at://xn--ugbaf6g.example/?#",
			ExpectedAuthority: "xn--ugbaf6g.example",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/?#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
		},



		{
			URI:          "at://localhost/com.example.fooBar",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
		},
		{
			URI:          "at://example.com/com.example.fooBar",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://example.com./com.example.fooBar",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
		},



		{
			URI:          "at://localhost/com.example.fooBar/",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
		},
		{
			URI:          "at://example.com/com.example.fooBar/",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://example.com./com.example.fooBar/",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
		},



		{
			URI:          "at://localhost/com.example.fooBar?",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
		},
		{
			URI:          "at://example.com/com.example.fooBar?",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://example.com./com.example.fooBar?",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar?",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/?",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar?",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar?",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
		},



		{
			URI:          "at://localhost/com.example.fooBar#",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
		},
		{
			URI:          "at://example.com/com.example.fooBar#",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://example.com./com.example.fooBar#",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar#",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar#",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar#",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
		},



		{
			URI:          "at://localhost/com.example.fooBar/?#",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
		},
		{
			URI:          "at://example.com/com.example.fooBar/?#",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://example.com./com.example.fooBar/?#",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/?#",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/?#",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/?#",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/?#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
		},



		{
			URI:          "at://localhost/com.example.fooBar/3jui7kd54zh2y",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
			ExpectedRKey:                                    "3jui7kd54zh2y",
		},
		{
			URI:          "at://example.com/com.example.fooBar/3jui7kd54zh2y",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                      "3jui7kd54zh2y",
		},
		{
			URI:          "at://example.com./com.example.fooBar/3jui7kd54zh2y",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
			ExpectedRKey:                                       "3jui7kd54zh2y",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/3jui7kd54zh2y",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/3jui7kd54zh2y",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                     "3jui7kd54zh2y",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/3jui7kd54zh2y",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/3jui7kd54zh2y",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
			ExpectedRKey:                                                           "3jui7kd54zh2y",
		},



		{
			URI:          "at://localhost/com.example.fooBar/3jui7kd54zh2y?",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
			ExpectedRKey:                                    "3jui7kd54zh2y",
		},
		{
			URI:          "at://example.com/com.example.fooBar/3jui7kd54zh2y?",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                      "3jui7kd54zh2y",
		},
		{
			URI:          "at://example.com./com.example.fooBar/3jui7kd54zh2y?",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
			ExpectedRKey:                                       "3jui7kd54zh2y",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/3jui7kd54zh2y?",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/3jui7kd54zh2y?",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                     "3jui7kd54zh2y",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/3jui7kd54zh2y?",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/3jui7kd54zh2y?",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
			ExpectedRKey:                                                           "3jui7kd54zh2y",
		},



		{
			URI:          "at://localhost/com.example.fooBar/3jui7kd54zh2y#",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
			ExpectedRKey:                                    "3jui7kd54zh2y",
		},
		{
			URI:          "at://example.com/com.example.fooBar/3jui7kd54zh2y#",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                      "3jui7kd54zh2y",
		},
		{
			URI:          "at://example.com./com.example.fooBar/3jui7kd54zh2y#",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
			ExpectedRKey:                                       "3jui7kd54zh2y",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/3jui7kd54zh2y#",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/3jui7kd54zh2y#",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                     "3jui7kd54zh2y",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/3jui7kd54zh2y#",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/3jui7kd54zh2y#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
			ExpectedRKey:                                                           "3jui7kd54zh2y",
		},



		{
			URI:          "at://localhost/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
			ExpectedRKey:                                    "3jui7kd54zh2y",
			ExpectedQuery:                                                 "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://example.com/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                      "3jui7kd54zh2y",
			ExpectedQuery:                                                   "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://example.com./com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
			ExpectedRKey:                                       "3jui7kd54zh2y",
			ExpectedQuery:                                                    "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
			ExpectedQuery:                                                           "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                     "3jui7kd54zh2y",
			ExpectedQuery:                                                  "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
			ExpectedQuery:                                                           "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
			ExpectedRKey:                                                           "3jui7kd54zh2y",
			ExpectedQuery:                                                                        "once=1&twice=2&thrice=3&fource=4",
		},



		{
			URI:          "at://localhost/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
			ExpectedRKey:                                    "3jui7kd54zh2y",
			ExpectedQuery:                                                 "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://example.com/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                      "3jui7kd54zh2y",
			ExpectedQuery:                                                   "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://example.com./com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
			ExpectedRKey:                                       "3jui7kd54zh2y",
			ExpectedQuery:                                                    "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
			ExpectedQuery:                                                           "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                     "3jui7kd54zh2y",
			ExpectedQuery:                                                  "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
			ExpectedQuery:                                                           "once=1&twice=2&thrice=3&fource=4",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
			ExpectedRKey:                                                           "3jui7kd54zh2y",
			ExpectedQuery:                                                                        "once=1&twice=2&thrice=3&fource=4",
		},



		{
			URI:          "at://localhost/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			ExpectedAuthority: "localhost",
			ExpectedCollection:          "com.example.fooBar",
			ExpectedRKey:                                    "3jui7kd54zh2y",
			ExpectedQuery:                                                 "once=1&twice=2&thrice=3&fource=4",
			ExpectedFragment:                                                                               "path(/apple/banana/cherry)",
		},
		{
			URI:          "at://example.com/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			ExpectedAuthority: "example.com",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                      "3jui7kd54zh2y",
			ExpectedQuery:                                                   "once=1&twice=2&thrice=3&fource=4",
			ExpectedFragment:                                                                                 "path(/apple/banana/cherry)",
		},
		{
			URI:          "at://example.com./com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			ExpectedAuthority: "example.com.",
			ExpectedCollection:             "com.example.fooBar",
			ExpectedRKey:                                       "3jui7kd54zh2y",
			ExpectedQuery:                                                    "once=1&twice=2&thrice=3&fource=4",
			ExpectedFragment:                                                                                  "path(/apple/banana/cherry)",
		},
		{
			URI:          "at://apple.banana.cherry/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			ExpectedAuthority: "apple.banana.cherry",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
			ExpectedQuery:                                                           "once=1&twice=2&thrice=3&fource=4",
			ExpectedFragment:                                                                                         "path(/apple/banana/cherry)",
		},
		{
			URI:          "at://Hello.WORLD/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			ExpectedAuthority: "hello.world",
			ExpectedCollection:            "com.example.fooBar",
			ExpectedRKey:                                     "3jui7kd54zh2y",
			ExpectedQuery:                                                  "once=1&twice=2&thrice=3&fource=4",
			ExpectedFragment:                                                                                "path(/apple/banana/cherry)",
		},
		{
			URI:          "at://xn--ugbaf6g.example/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			ExpectedAuthority: "xn--ugbaf6g.example",
			ExpectedCollection:                    "com.example.fooBar",
			ExpectedRKey:                                              "3jui7kd54zh2y",
			ExpectedQuery:                                                           "once=1&twice=2&thrice=3&fource=4",
			ExpectedFragment:                                                                                         "path(/apple/banana/cherry)",
		},
		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/3jui7kd54zh2y?once=1&twice=2&thrice=3&fource=4#path(/apple/banana/cherry)",
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
			ExpectedRKey:                                                           "3jui7kd54zh2y",
			ExpectedQuery:                                                                        "once=1&twice=2&thrice=3&fource=4",
			ExpectedFragment:                                                                                                      "path(/apple/banana/cherry)",
		},



		{
			URI:          "at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/" + strings.Repeat("0123456789ABCDEFGHIJKLMNOPQRSTUV", 256)[len("at://did:plc:scewmn2pl3oz36mxme2b6czz/com.example.fooBar/"):],
			ExpectedAuthority: "did:plc:scewmn2pl3oz36mxme2b6czz",
			ExpectedCollection:                                 "com.example.fooBar",
			ExpectedRKey:                                                           "PQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV0123456789ABCDEFGHIJKLMNOPQRSTUV",
		},
	}

	for testNumber, test := range tests {

		actualAuthority, actualCollection, actualRKey, actualQuery, actualFragment, err := aturi.SplitAndNormalizeAndValidate(test.URI)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("URI: %q", test.URI)
			continue
		}

		{
			expected := test.ExpectedAuthority
			actual   :=        actualAuthority

			if expected != actual {
				t.Errorf("For test #%d, the actual 'authority' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URI: %q", test.URI)
				continue
			}
		}

		{
			expected := test.ExpectedCollection
			actual   :=        actualCollection

			if expected != actual {
				t.Errorf("For test #%d, the actual 'collection' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URI: %q", test.URI)
				continue
			}
		}

		{
			expected := test.ExpectedRKey
			actual   :=        actualRKey

			if expected != actual {
				t.Errorf("For test #%d, the actual 'rkey' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URI: %q", test.URI)
				continue
			}
		}

		{
			expected := test.ExpectedQuery
			actual   :=        actualQuery

			if expected != actual {
				t.Errorf("For test #%d, the actual 'query' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URI: %q", test.URI)
				continue
			}
		}

		{
			expected := test.ExpectedFragment
			actual   :=        actualFragment

			if expected != actual {
				t.Errorf("For test #%d, the actual 'fragment' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URI: %q", test.URI)
				continue
			}
		}

	}
}

func TestSplitAndNormalizeAndValidate_fail(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: `aturi: empty URI`,
		},


		{
			URI: "apple",
			ExpectedError: `aturi: URI "apple" is not an AT-URI because it does not begin with "at:"`,
		},
		{
			URI: "banana",
			ExpectedError: `aturi: URI "banana" is not an AT-URI because it does not begin with "at:"`,
		},
		{
			URI: "cherry",
			ExpectedError: `aturi: URI "cherry" is not an AT-URI because it does not begin with "at:"`,
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
			URI:                          "at://user:pass@foo.com",
			ExpectedError: `aturi: AT-URI "at://user:pass@foo.com" may not have an "@" in its authority "user:pass@foo.com"`,
		},
		{
			URI:                          "at://user:pass@Foo.COM",
			ExpectedError: `aturi: AT-URI "at://user:pass@Foo.COM" may not have an "@" in its authority "user:pass@foo.com"`,
		},
		{
			URI:                          "at://@",
			ExpectedError: `aturi: AT-URI "at://@" may not have an "@" in its authority "@"`,
		},
		{
			URI:                          "at://@example",
			ExpectedError: `aturi: AT-URI "at://@example" may not have an "@" in its authority "@example"`,
		},
		{
			URI:                          "at://@example.com",
			ExpectedError: `aturi: AT-URI "at://@example.com" may not have an "@" in its authority "@example.com"`,
		},
	}

	for testNumber, test := range tests {

		_, _, _, _, _, err := aturi.SplitAndNormalizeAndValidate(test.URI)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("URI: %q", test.URI)
			continue
		}

		{
			expected := test.ExpectedError
			actual := err.Error()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'error' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("URI: %q", test.URI)
				continue
			}
		}
	}
}
