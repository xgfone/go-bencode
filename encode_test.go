package bencode_test

import (
	"testing"

	bencode "github.com/xgfone/go-bencode"
)

func TestBEncode(t *testing.T) {
	// dd := make(bencode.Dict)
	// dd["t"] = "aa"
	// dd["y"] = "r"

	// ll := make(bencode.List, 0)
	// ll = append(ll, "axje.u", "idhtnm")

	// aa := make(bencode.Dict)
	// aa["id"] = "abcdefghij0123456789"
	// aa["token"] = "aoeusnth"
	// aa["values"] = ll

	// dd["r"] = aa

	dd := bencode.Dict{
		"t": "aa",
		"y": "r",
		"r": bencode.Dict{
			"id":     "abcdefghij0123456789",
			"token":  "aoeusnth",
			"values": bencode.List{"axje.u", "idhtnm"},
		},
	}

	result := "d1:rd2:id20:abcdefghij01234567895:token8:aoeusnth6:valuesl6:axje.u6:idhtnmee1:t2:aa1:y1:re"
	_r := bencode.BEncode(dd)
	if result != string(_r) {
		t.Fail()
	}
}
