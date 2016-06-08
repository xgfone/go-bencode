package bencode

import (
	"bytes"
	"fmt"
	"sort"
)

type bencoder struct{}

func NewBEncode() IBEncode {
	return bencoder{}
}

func (b bencoder) Int64(v int64) []byte {
	return []byte(fmt.Sprintf("i%de", v))
}

func (b bencoder) String(v string) []byte {
	return []byte(fmt.Sprintf("%d:%s", len(v), v))
}

func (b bencoder) List(v List) []byte {
	var buf bytes.Buffer
	var tmp []byte
	for _, e := range v {
		switch e.(type) {
		case int:
			tmp = b.Int64(int64(e.(int)))
		case int64:
			tmp = b.Int64(e.(int64))
		case string:
			tmp = b.String(e.(string))
		case List:
			tmp = b.List(e.(List))
		case Dict:
			tmp = b.Dict(e.(Dict))
		default:
			return nil
		}

		if tmp == nil {
			return nil
		}

		buf.Write(tmp)
	}
	return []byte(fmt.Sprintf("l%se", buf.Bytes()))
}

func (b bencoder) Dict(v Dict) []byte {
	keys := getKeyList(v)
	sort.Strings(keys)
	var buf bytes.Buffer
	var _key []byte
	var _value []byte
	for _, key := range keys {
		_key = b.String(key)
		value, _ := v[key]
		switch value.(type) {
		case int:
			_value = b.Int64(int64(value.(int)))
		case int64:
			_value = b.Int64(value.(int64))
		case string:
			_value = b.String(value.(string))
		case List:
			_value = b.List(value.(List))
		case Dict:
			_value = b.Dict(value.(Dict))
		default:
			return nil
		}

		if _value == nil {
			return nil
		}

		// buf.Write([]byte(fmt.Sprintf("%s%s", _key, _value)))
		buf.Write(_key)
		buf.Write(_value)
	}
	return []byte(fmt.Sprintf("d%se", buf.Bytes()))
}

func getKeyList(d Dict) (keys []string) {
	for key, _ := range d {
		keys = append(keys, key)
	}
	return
}
