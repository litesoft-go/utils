package opaquestrings

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/codec"
)

//noinspection GoUnusedExportedFunction
func GetOpaqueField(curOpaque, key string) (value string, err error) {
	var keyValues map[string]string
	if (curOpaque != "") && (key != "") {
		keyValues, err = decodeOpaqueState(curOpaque)
		value = keyValues[key]
	}
	return
}

//noinspection GoUnusedExportedFunction
func UpdateOpaqueField(curOpaque, key, newValue string, setter func(string)) (updated bool, err error) {
	if key == "" {
		return false, nil
	}
	var keyValues map[string]string
	keyValues, err = decodeOpaqueState(curOpaque)
	if err != nil {
		return false, err
	}
	curValue := keyValues[key]
	if curValue == newValue { // catches (!exists && (newValue == nil))
		return false, nil
	}
	if newValue == "" { // Must have existed!
		delete(keyValues, key)
	} else {
		if keyValues == nil {
			keyValues = make(map[string]string, 1)
		}
		keyValues[key] = newValue
	}
	var newOpaque string
	if len(keyValues) != 0 {
		newOpaque, err = codec.Encode(keyValues)
		if err != nil {
			return false, err
		}
	}
	setter(newOpaque)
	return true, nil
}

func decodeOpaqueState(opaqueString string) (keyValues map[string]string, err error) {
	if opaqueString != "" {
		keyValues, err = codec.Decode(opaqueString)
	}
	return
}
