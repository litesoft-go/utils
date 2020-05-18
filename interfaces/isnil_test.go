package interfaces

import (
	"testing"
	// standard libs only above!
)

type anInterface interface {
	sync()
}

type anImpl struct{}

func (in *anImpl) sync() {}

var (
	zString    string
	zStringPtr *string

	vChan    = make(chan int)
	vChanNil chan int

	vFunc    = func() string { return "func" }
	vFuncNil func()

	vInterface    anInterface = &anImpl{}
	vInterfaceNil interface{}

	zMap    = map[string]string{}
	zMapNil map[string]string // Map

	zSlice    = make([]string, 0)
	zSliceNil []string

	zArray    = [2]string{}
	zArrayPtr *[2]string // Arrays are not 'special' like Chan, Func, Interface, Map, or Slice
)

func TestIsNil(t *testing.T) {
	checkNil(t, "zString", zString, zStringPtr)
	checkNil(t, "vChan", vChan, vChanNil)
	checkNil(t, "vFunc", vFunc, vFuncNil)
	checkNil(t, "vInterface", vInterface, vInterfaceNil)
	checkNil(t, "zMap", zMap, zMapNil)
	checkNil(t, "zSlice", zSlice, zSliceNil)
	checkNil(t, "zArray", zArray, zArrayPtr)
}

func checkNil(t *testing.T, what string, regValue, nilValue interface{}) {
	cnve(t, what, "reg", regValue, false)
	cnve(t, what, "nil", nilValue, true)
}

func cnve(t *testing.T, what, form string, value interface{}, expected bool) {
	isNil := IsNil(value)
	if isNil != expected {
		t.Errorf("%s %s IsNil reported: %v\n", what, form, isNil)
	}
}
