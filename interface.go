package bencode

import "errors"

var (
	FMTERR = errors.New("The bencode format is error")
)

type List []interface{}
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
	case int64:
		return _bencoder.Int64(v.(int64))
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
