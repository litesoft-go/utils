package codec

import (
	"errors"
	"strings"
	// standard libs only above!
)

var id = ID_ASCI

// ASCII (Raw) codec - limitations: No ';' in any key's value; and as usual keys may not contain '='!
var ASCII Codec

func init() {
	ASCII = asciiCodec{}
	Register(id, ASCII)
}

type asciiCodec struct{}

var idAsPrefix = id + ":"

func (r asciiCodec) Encode(keyValues map[string]string) (opaque string, err error) {
	if len(keyValues) == 0 {
		opaque = idAsPrefix
		return
	}
	collector := make([]string, (len(keyValues)*4)+1)
	collector = append(collector, idAsPrefix)
	prefix := ""
	for k, v := range keyValues {
		k, v, err = normalize(k, v)
		if err != nil {
			return
		}
		if v != "" {
			collector = append(collector, prefix, k, "=", v)
			prefix = ";"
		}
	}
	opaque = strings.Join(collector, "")
	return
}

func (r asciiCodec) Decode(opaque string) (keyValues map[string]string, err error) {
	if !strings.HasPrefix(opaque, idAsPrefix) {
		err = errors.New("codec ID, expected '" + id + "', but NOT at front of: " + opaque)
		return
	}
	data := string([]byte(opaque)[len(idAsPrefix):])
	if data != "" {
		var k, v string
		pairs := strings.Split(data, ";")
		keyValues = make(map[string]string, len(pairs))
		for _, pair := range pairs {
			k, v, err = normalizePair(pair)
			if err != nil {
				return
			}
			if v != "" {
				keyValues[k] = v
			}
		}
	}
	return
}

func normalizePair(pair string) (key, value string, err error) {
	pair = strings.TrimSpace(pair)
	if pair != "" { // ignore empty pair?
		parts := strings.Split(pair, "=")
		switch len(parts) {
		case 1:
			err = errors.New("a key/value pair did not contain a '=', pair was: " + pair)
		case 2:
			key, value, err = normalize(parts[0], parts[1])
		default:
			err = errors.New("key in key/value pair contained '='(s), pair was: " + pair)
		}
	}
	return
}

func normalize(k, v string) (key, value string, err error) {
	key = strings.TrimSpace(k)
	value = strings.TrimSpace(v)
	switch {
	case key == "":
		err = errors.New("codecs do not allow empty keys; value was: " + value)
	case strings.Contains(key, "="):
		err = errors.New("codecs do not allow '=' in keys; key was '" + key + "'")
	case strings.Contains(value, ";"):
		err = errors.New("codec (Ascii) does not allow ';' in values; key '" + key + "' value was: " + value)
	}
	return
}
