package opaquestrings

import (
	"strconv"
	"time"

	// standard libs only above!

	"github.com/litesoft-go/utils/int32s"
	"github.com/litesoft-go/utils/iso8601"
	"github.com/litesoft-go/utils/strs"
)

func (p *OpaqueProxy) GetStringP(key string) (v *string, err error) {
	var value string
	var exists bool
	value, exists, err = p.RawGetString(key)
	if exists {
		v = &value
	}
	return
}

func (p *OpaqueProxy) GetString(key string) (value string, err error) {
	value, _, err = p.RawGetString(key)
	return
}

func (p *OpaqueProxy) UpdateString(key, newValue string) (updated bool, err error) {
	return p.RawUpdateString(key, newValue)
}

func (p *OpaqueProxy) UpdateStringP(key string, newValue *string) (updated bool, err error) {
	return p.UpdateString(key, strs.FromOptional(newValue, ""))
}

func (p *OpaqueProxy) GetInt32P(key string) (v *int32, err error) {
	var value string
	var exists bool
	value, exists, err = p.RawGetString(key)
	if exists {
		v, err = int32s.OptionalFromA(value) // v = nil if err !nil
	}
	return
}

func (p *OpaqueProxy) GetInt32(key string) (v int32, err error) {
	var value *int32
	value, err = p.GetInt32P(key)
	if value != nil {
		v = *value
	}
	return
}

func (p *OpaqueProxy) UpdateInt32(key string, newValue int32) (updated bool, err error) {
	return p.UpdateString(key, int32s.ToA(newValue))
}

func (p *OpaqueProxy) UpdateInt32P(key string, newValue *int32) (updated bool, err error) {
	return p.UpdateString(key, int32s.OptionalToA(newValue, ""))
}

func (p *OpaqueProxy) GetTimeP(key string) (v *time.Time, err error) {
	var value string
	var exists bool
	value, exists, err = p.RawGetString(key)
	if exists {
		var t time.Time
		t, err = iso8601.ZmillisFromString(value)
		if err == nil {
			v = &t
		}
	}
	return
}

func (p *OpaqueProxy) GetTime(key string) (v time.Time, err error) {
	var value *time.Time
	value, err = p.GetTimeP(key)
	if value != nil {
		v = *value
	}
	return
}

func (p *OpaqueProxy) UpdateTime(key string, newValue time.Time) (updated bool, err error) {
	return p.UpdateString(key, iso8601.ZmillisToString(&newValue))
}

func (p *OpaqueProxy) UpdateTimeP(key string, newValue *time.Time) (updated bool, err error) {
	return p.UpdateString(key, iso8601.OptionalToA(newValue, ""))
}

func (p *OpaqueProxy) GetBoolP(key string) (v *bool, err error) {
	var value string
	var exists bool
	value, exists, err = p.RawGetString(key)
	if exists {
		var b bool
		b, err = strconv.ParseBool(value)
		if err == nil {
			v = &b
		}
	}
	return
}

func (p *OpaqueProxy) GetBool(key string) (v bool, err error) {
	var value *bool
	value, err = p.GetBoolP(key)
	if value != nil {
		v = *value
	}
	return
}

func (p *OpaqueProxy) UpdateBool(key string, newValue bool) (updated bool, err error) {
	return p.UpdateString(key, strconv.FormatBool(newValue))
}

func (p *OpaqueProxy) UpdateBoolP(key string, newValue *bool) (updated bool, err error) {
	if newValue != nil {
		return p.UpdateBool(key, *newValue)
	}
	return p.Delete(key)
}
