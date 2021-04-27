package simpleconf

import (
	"fmt"
	"github.com/litesoft-go/utils/ints"
	"os"
	"strings"
)

type ArgProxy struct {
	argsWithProg []string
	proxyTo      Config
}

//noinspection GoUnusedExportedFunction
func WithArgs(args []string) (config *ArgProxy) {
	config = &ArgProxy{argsWithProg: args}
	return
}

func (in *ArgProxy) ValuesFor(key string) (values []string) {
	found, plus, equals := in.fromArgs(key)
	if equals {
		return found
	}
	values = in.proxyTo.ValuesFor(key)
	if plus {
		values = append(values, found...)
	}
	return
}

func (in *ArgProxy) ValueOf(key string) (value string, err error) {
	return in.ExtractValueOf(key, in.ValuesFor(key))
}

func (in *ArgProxy) ExtractValueOf(key string, from []string) (value string, err error) {
	return in.proxyTo.ExtractValueOf(key, from)
}

//noinspection GoUnusedExportedFunction
func (in *ArgProxy) Load(proxyTo Config, errIn error) (config *ArgProxy, err error) {
	in.proxyTo = proxyTo
	return in, errIn
}

//noinspection GoUnusedExportedFunction
func (in *ArgProxy) CheckNoMoreArgs() {
	if len(in.argsWithProg) > 1 {
		fmt.Println("No additional CLI parameters supported, but got: ", in.argsWithProg[1:])
		os.Exit(1)
	}
}

func (in *ArgProxy) fromArgs(key string) (values []string, plus, equals bool) {
	if len(key) != 0 {
		i := 1
		args := in.argsWithProg
		for i < len(args) {
			entry := args[i]
			if strings.HasPrefix(entry, "-") {
				dashes := ints.Tertiary(strings.HasPrefix(entry, "--"), 2, 1)
				entry = entry[dashes:]
				if strings.HasPrefix(entry, key) {
					entry = entry[len(key):]
					if len(entry) > 1 {
						firstChar := entry[0]
						if (firstChar == '=') || (firstChar == '+') {
							values = append(values, entry[1:])
							if !equals {
								if firstChar == '=' {
									equals = true
								} else {
									plus = true
								}
							}
							args = append(args[:i], args[i+1:]...)
							continue
						}
					}
				}
			}
			i++
		}
		in.argsWithProg = args
	}
	return
}
