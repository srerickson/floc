package ocfl

import (
	"encoding/json"
	"testing"
)

func TestCueValidate(t *testing.T) {
	inv := &Inventory{}
	test := `{
		"digestAlgorithm": "sha512",
		"fixity": {
		"md5": {
			"184f84e28cbe75e050e9c25ea7f2e939": [ "v1/content/foo/bar.xml" ],
			"2673a7b11a70bc7ff960ad8127b4adeb": [ "v2/content/foo/bar.xml" ],
			"c289c8ccd4bab6e385f5afdd89b5bda2": [ "v1/content/image.tiff" ],
			"d41d8cd98f00b204e9800998ecf8427e": [ "v1/content/empty.txt" ]
		},
		"sha1": {
			"66709b068a2faead97113559db78ccd44712cbf2": [ "v1/content/foo/bar.xml" ],
			"a6357c99ecc5752931e133227581e914968f3b9c": [ "v2/content/foo/bar.xml" ],
			"b9c7ccc6154974288132b63c15db8d2750716b49": [ "v1/content/image.tiff" ],
			"da39a3ee5e6b4b0d3255bfef95601890afd80709": [ "v1/content/empty.txt" ]
		}
		},
		"head": "v3",
		"id": "ark:/12345/bcd987",
		"manifest": {
		"4d27c8b531": [ "v2/content/foo/bar.xml" ],
		"7dcc35c311": [ "v1/content/foo/bar.xml" ],
		"cf83e1a3e1": [ "v1/content/empty.txt" ],
		"ffccf662e1": [ "v1/content/image.tiff" ]
		},
		"type": "https://ocfl.io/1.0/spec/#inventory",
		"versions": {
		"v1": {
			"created": "2018-01-01T01:01:01Z",
			"message": "Initial import",
			"state": {
			"17dcc35c31": [ "foo/bar.xml" ],
			"1cf83e1a3e": [ "empty.txt" ],
			"1ffccf662e": [ "image.tiff" ]
			},
			"user": {
			"address": "alice@example.com",
			"name": "Alice"
			}
		},
		"v2": {
			"created": "2018-02-02T02:02:02Z",
			"message": "Fix bar.xml, remove image.tiff, add empty2.txt",
			"state": {
			"4d27c8b531": [ "foo/bar.xml" ],
			"cf83e1a3e1": [ "empty.txt", "empty2.txt" ]
			},
			"user": {
			"address": "bob@example.com",
			"name": "Bob"
			}
		},
		"v3": {
			"created": "2018-03-03T03:03:03Z",
			"message": "Reinstate image.tiff, delete empty.txt",
			"state": {
			"4d27c8b53a": [ "foo/bar.xml" ],
			"cf83e1a3ea": [ "empty2.txt" ],
			"ffccf662ea": [ "image.tiff" ]
			},
			"user": {
			"address": "cecilia@example.com",
			"name": "Cecilia"
			}
		}
		}
	}`
	if err := json.Unmarshal([]byte(test), inv); err != nil {
		t.Error(err)
	}

	err := inv.CUEValudate()
	if err != nil {
		t.Error(err)
	}
}