package bencode_test

import (
	"testing"

	bencode "github.com/xgfone/go-bencode"
)

func TestBDecode(t *testing.T) {
	buf := []byte("d1:rd2:id20:abcdefghij01234567895:token8:aoeusnth6:valuesl6:axje.u6:idhtnmee1:t2:aa1:y1:re")
	if _, err := bencode.BDecode(buf); err != nil {
		t.Fail()
	}
}
