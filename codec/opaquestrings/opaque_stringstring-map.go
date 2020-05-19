package opaquestrings

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/codec"
	"github.com/litesoft-go/utils/strs"
)

//noinspection GoUnusedExportedFunction
func GetOpaqueField(curOpaque, key string) (value string, exists bool, err error) {
	if (curOpaque != "") && (key != "") {
		var keyValues map[string]string
		keyValues, err = decodeOpaqueState(curOpaque)
		if err == nil {
			value, exists = keyValues[key]
		}
	}
	return
}

//noinspection GoUnusedExportedFunction
func UpdateOpaqueField(curOpaque, key, newValue string, setter strs.Sync) (updated bool, err error) {
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

//noinspection GoUnusedExportedFunction
func DeleteOpaqueField(curOpaque, key string, setter strs.Sync) (updated bool, err error) {
	return UpdateOpaqueField(curOpaque, key, "", setter)
}

func decodeOpaqueState(opaqueString string) (keyValues map[string]string, err error) {
	if opaqueString != "" {
		keyValues, err = codec.Decode(opaqueString)
	}
	return
}
