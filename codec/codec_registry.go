package codec

import (
	"errors"
	"strings"
)

var registry = make(map[string]Codec, 1)
var defaultEncoder string

//noinspection GoUnusedExportedFunction
func SetDefaultEncoder(id string) {
	defaultEncoder = id
}

func Register(id string, codec Codec) {
	existing, exists := registry[id]
	if !exists {
		registry[id] = codec
		if id > defaultEncoder {
			defaultEncoder = id
		}
		return
	}
	if existing != codec {
		panic("Duplicate 'codec' registration for '" + id + "'")
	}
}

func Encode(keyValues map[string]string) (opaque string, err error) {
	codec, ok := registry[defaultEncoder]
	if ok {
		return codec.Encode(keyValues)
	}
	panic("no default Encoder")
}

func Decode(opaque string) (keyValues map[string]string, err error) {
	index := strings.Index(opaque, SeparatorBetweenContentAndID)
	if index < 1 {
		err = errors.New("no codec ID at front of: " + opaque)
		return
	}
	id := string([]byte(opaque)[:index])
	codec, ok := registry[id]
	if ok {
		return codec.Decode(opaque)
	}
	err = errors.New("no codec registered for ID '" + id + "' from: " + opaque)
	return
}
