package bencode

import "errors"

var (
	FMTERR = errors.New("The bencode format is error")
)

// The element of List may be integer, string, List, Dict.
// Integer is one of int, int8, int16, int32, int64, uint, uint8, uint16,
// uint32, uint64. When decoding the data, integer is int64.
type List []interface{}

// The element of List may be integer, string, List, Dict.
// Integer is one of int, int8, int16, int32, int64, uint, uint8, uint16,
// uint32, uint64. When decoding the data, integer is int64.
type Dict map[string]interface{}

type IBEncode interface {
	Int64(v int64) []byte
	String(v string) []byte
	List(v List) []byte
	Dict(v Dict) []byte
}

type IBDecode interface {
	Int64(buf []byte) (int64, error)
	String(buf []byte) (string, error)
	List(buf []byte) (List, error)
	Dict(buf []byte) (Dict, error)
}

var _bencoder bencoder

func BEncode(v interface{}) []byte {
	switch v.(type) {
	case int:
		return _bencoder.Int64(int64(v.(int)))
	case int8:
		return _bencoder.Int64(int64(v.(int8)))
	case int16:
		return _bencoder.Int64(int64(v.(int16)))
	case int32:
		return _bencoder.Int64(int64(v.(int32)))
	case int64:
		return _bencoder.Int64(v.(int64))

	case uint:
		return _bencoder.Int64(int64(v.(uint)))
	case uint8:
		return _bencoder.Int64(int64(v.(uint8)))
	case uint16:
		return _bencoder.Int64(int64(v.(uint16)))
	case uint32:
		return _bencoder.Int64(int64(v.(uint32)))
	case uint64:
		return _bencoder.Int64(int64(v.(uint64)))

	case string:
		return _bencoder.String(v.(string))
	case List:
		return _bencoder.List(v.(List))
	case Dict:
		return _bencoder.Dict(v.(Dict))
	default:
		return nil
	}
}

var _bdecoder bdecoder

func BDecode(buf []byte) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	switch buf[0] {
	case 'i':
		result, err = _bdecoder.Int64(buf)
	case 'l':
		result, err = _bdecoder.List(buf)
	case 'd':
		result, err = _bdecoder.Dict(buf)
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		result, err = _bdecoder.String(buf)
	}

	return result, err
}
