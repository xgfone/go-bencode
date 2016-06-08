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

// Encode the value of int64.
func (b bencoder) Int64(v int64) []byte {
	return []byte(fmt.Sprintf("i%de", v))
}

// Encode the value of string.
func (b bencoder) String(v string) []byte {
	return []byte(fmt.Sprintf("%d:%s", len(v), v))
}

// Encode the value of List.
func (b bencoder) List(v List) []byte {
	var buf bytes.Buffer
	var tmp []byte
	for _, e := range v {
		switch e.(type) {
		case int8:
			tmp = b.Int64(int64(e.(int8)))
		case int16:
			tmp = b.Int64(int64(e.(int16)))
		case int32:
			tmp = b.Int64(int64(e.(int32)))
		case int64:
			tmp = b.Int64(e.(int64))
		case int:
			tmp = b.Int64(int64(e.(int)))
		case uint8:
			tmp = b.Int64(int64(e.(uint8)))
		case uint16:
			tmp = b.Int64(int64(e.(uint16)))
		case uint32:
			tmp = b.Int64(int64(e.(uint32)))
		case uint64:
			tmp = b.Int64(int64(e.(uint64)))
		case uint:
			tmp = b.Int64(int64(e.(uint)))
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

// Encode the value of Dict.
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
		case int8:
			_value = b.Int64(int64(value.(int8)))
		case int16:
			_value = b.Int64(int64(value.(int16)))
		case int32:
			_value = b.Int64(int64(value.(int32)))
		case int64:
			_value = b.Int64(value.(int64))

		case uint:
			_value = b.Int64(int64(value.(uint)))
		case uint8:
			_value = b.Int64(int64(value.(uint8)))
		case uint16:
			_value = b.Int64(int64(value.(uint16)))
		case uint32:
			_value = b.Int64(int64(value.(uint32)))
		case uint64:
			_value = b.Int64(int64(value.(uint64)))

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
