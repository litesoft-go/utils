package opaquestrings

import "github.com/litesoft-go/utils/strs"

type OpaqueProxy struct {
	Get strs.Source
	Put strs.Sync
}

// RawGetString will fetch a field by key, as a String, from the OpaqueProxy.
// If err is !nil then exists will be false!
//noinspection GoUnusedExportedFunction
func (p *OpaqueProxy) RawGetString(key string) (value string, exists bool, err error) {
	return GetOpaqueField(p.Get(), key)
}

// RawUpdateString will set a field, as a String, if it is different then the current value, via the OpaqueProxy.
// Iff err is !nil can updated be true!
//noinspection GoUnusedExportedFunction
func (p *OpaqueProxy) RawUpdateString(key, newValue string) (updated bool, err error) {
	return UpdateOpaqueField(p.Get(), key, newValue, p.Put)
}

func (p *OpaqueProxy) Delete(key string) (updated bool, err error) {
	return DeleteOpaqueField(p.Get(), key, p.Put)
}
