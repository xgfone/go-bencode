package bencode

import (
	"bytes"
	"strconv"
)

type bdecoder struct{}

func NewBDecode() IBDecode {
	return bdecoder{}
}

func (b bdecoder) Int64(buf []byte) (v int64, err error) {
	v, _, err = b.toInt64(buf)
	return
}

func (b bdecoder) String(buf []byte) (v string, err error) {
	v, _, err = b.toString(buf)
	return
}

func (b bdecoder) List(buf []byte) (v List, err error) {
	v, _, err = b.toList(buf)
	return
}

func (b bdecoder) Dict(buf []byte) (v Dict, err error) {
	v, _, err = b.toDict(buf)
	return
}

func (b bdecoder) toInt64(buf []byte) (int64, []byte, error) {
	if buf == nil || buf[0] != 'i' {
		return 0, buf, FMTERR
	}

	eindex := bytes.IndexByte(buf, 'e')
	if eindex == -1 || eindex == 1 {
		return 0, buf, FMTERR
	}

	if i, err := strconv.ParseInt(string(buf[1:eindex]), 10, 64); err != nil {
		return 0, buf, FMTERR
	} else {
		return i, buf[eindex+1:], nil
	}
}

func (b bdecoder) toString(buf []byte) (string, []byte, error) {
	if buf == nil {
		return "", buf, FMTERR
	}

	cindex := bytes.IndexByte(buf, ':')
	if cindex == -1 {
		return "", buf, FMTERR
	}

	lenght, err := strconv.Atoi(string(buf[:cindex]))
	if err != nil {
		return "", buf, FMTERR
	}

	start := cindex + 1
	end := start + lenght
	return string(buf[start:end]), buf[end:], nil
}

func (b bdecoder) toList(buf []byte) (List, []byte, error) {
	if buf == nil || buf[0] != 'l' || len(buf) <= 1 {
		return nil, buf, FMTERR
	}
	orgi_buf := buf
	buf = buf[1:]
	result := make(List, 0)

	for len(buf) > 0 {
		if buf[0] == 'e' {
			return result, buf[1:], nil
		}

		switch buf[0] {
		case 'i':
			if i, _buf, err := b.toInt64(buf); err == nil {
				result = append(result, i)
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		case 'l':
			if list, _buf, err := b.toList(buf); err == nil {
				result = append(result, list)
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		case 'd':
			if dict, _buf, err := b.toDict(buf); err == nil {
				result = append(result, dict)
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			if str, _buf, err := b.toString(buf); err == nil {
				result = append(result, str)
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		default:
			return nil, orgi_buf, FMTERR
		}
	}
	return nil, orgi_buf, FMTERR
}

func (b bdecoder) toDict(buf []byte) (Dict, []byte, error) {
	if buf == nil || buf[0] != 'd' || len(buf) <= 1 {
		return nil, buf, FMTERR
	}
	orgi_buf := buf
	buf = buf[1:]
	dict := make(Dict)
	var err error
	var key string

	for len(buf) > 0 {
		if buf[0] == 'e' {
			return dict, buf[1:], nil
		}

		key, buf, err = b.toString(buf)
		if err != nil || len(buf) <= 1 {
			return nil, orgi_buf, err
		}

		switch buf[0] {
		case 'i':
			if i, _buf, err := b.toInt64(buf); err == nil {
				dict[key] = i
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		case 'l':
			if list, _buf, err := b.toList(buf); err == nil {
				dict[key] = list
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		case 'd':
			if _dict, _buf, err := b.toDict(buf); err == nil {
				dict[key] = _dict
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			if str, _buf, err := b.toString(buf); err == nil {
				dict[key] = str
				buf = _buf
			} else {
				return nil, orgi_buf, err
			}
		default:
			return nil, orgi_buf, FMTERR
		}
	}

	return nil, orgi_buf, FMTERR
}
