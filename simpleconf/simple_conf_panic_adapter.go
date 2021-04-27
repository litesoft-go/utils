package simpleconf

import (
	"fmt"
)

type PanicAdapter struct {
	proxyTo Config
}

func (in *PanicAdapter) ValueOfOptional(key string) string {
	value, err := in.proxyTo.ValueOf(key)
	if err != nil {
		panic(err)
	}
	return value
}

func (in *PanicAdapter) ValuesForOptional(key string) []string {
	return in.proxyTo.ValuesFor(key)
}

func (in *PanicAdapter) ValueOfRequired(key string) (value string) {
	value = in.ValueOfOptional(key)
	if len(value) == 0 {
		panic(fmt.Errorf("no value for '%s', but required", key))
	}
	return
}

func (in *PanicAdapter) ValuesForRequired(key string) (values []string) {
	values = in.ValuesForOptional(key)
	if len(values) == 0 {
		panic(fmt.Errorf("no values for '%s', but required", key))
	}
	return
}

//noinspection GoUnusedExportedFunction
func Adapter(proxyTo Config, err error) *PanicAdapter {
	if err != nil {
		panic(err)
	}
	return &PanicAdapter{proxyTo: proxyTo}
}
