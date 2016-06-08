# go-bencode
A Go implementation of bencode, see [BitTorrent Protocol](http://www.bittorrent.org/beps/bep_0003.html).

## Example
```go
package main

import (
    "fmt"

    bencode "github.com/xgfone/go-bencode"
)

func TestBEncode() {
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
        fmt.Println("TestBEncode ERROR")
    } else {
        fmt.Println("TestBEncode OK")
    }
}

func TestBDecode() {
    buf := []byte("d1:rd2:id20:abcdefghij01234567895:token8:aoeusnth6:valuesl6:axje.u6:idhtnmee1:t2:aa1:y1:re")
    if _dict, err := bencode.BDecode(buf); err != nil {
        fmt.Println("TestBDecode ERROR")
    } else {
        fmt.Println("TestBDecode OK")
        dict := _dict.(bencode.Dict)
        t, _ := dict["t"]
        fmt.Printf("t=%s\n", t.(string))

        y, _ := dict["y"]
        fmt.Printf("y=%s\n", y.(string))

        _r, _ := dict["r"]
        r := _r.(bencode.Dict)

        id, _ := r["id"]
        fmt.Printf("id=%s\n", id.(string))

        token, _ := r["token"]
        fmt.Printf("token=%s\n", token.(string))

        values, _ := r["values"]
        list := values.(bencode.List)

        fmt.Printf("list[%s, %s]\n", list[0].(string), list[1].(string))
    }
}

func main() {
    TestBEncode()
    TestBDecode()
}
```